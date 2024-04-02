package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jassue/gin-wire/app/service"
	"github.com/jassue/gin-wire/config"
	"github.com/jassue/gin-wire/util/path"
	"github.com/sony/sonyflake"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData, NewDB, NewRedis, NewMongoDB, NewTransaction,
	NewUserRepo, NewJwtRepo, NewMediaRepo,
	NewCompRepo, NewCAMsgRepo, NewPageRepo,
)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	sf  *sonyflake.Sonyflake
	mdb *mongo.Database
}

// NewData .
func NewData(logger *zap.Logger, db *gorm.DB, rdb *redis.Client, sf *sonyflake.Sonyflake, mdb *mongo.Database) (*Data, func(), error) {
	cleanup := func() {
		logger.Info("closing the data resources")
	}

	return &Data{db: db, rdb: rdb, sf: sf, mdb: mdb}, cleanup, nil
}

// 数据库连接

// NewDB .
func NewDB(conf *config.Configuration, gLog *zap.Logger) *gorm.DB {
	if conf.Database.Driver != "mysql" {
		panic(conf.Database.Driver + " driver is not supported")
	}

	var writer io.Writer
	var logMode logger.LogLevel

	// 是否启用日志文件
	if conf.Database.EnableFileLogWriter {
		logFileDir := conf.Log.RootDir
		if !filepath.IsAbs(logFileDir) {
			logFileDir = filepath.Join(path.RootPath(), logFileDir)
		}
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   filepath.Join(logFileDir, conf.Database.LogFilename),
			MaxSize:    conf.Log.MaxSize,
			MaxBackups: conf.Log.MaxBackups,
			MaxAge:     conf.Log.MaxAge,
			Compress:   conf.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}

	switch conf.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	newLogger := logger.New(
		log.New(writer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                        // 慢查询 SQL 阈值
			Colorful:                  !conf.Database.EnableFileLogWriter, // 禁用彩色打印
			IgnoreRecordNotFoundError: false,                              // 忽略ErrRecordNotFound（记录未找到）错误
			LogLevel:                  logMode,                            // Log lever
		},
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.Database.UserName,
		conf.Database.Password,
		conf.Database.Host,
		strconv.Itoa(conf.Database.Port),
		conf.Database.Database,
		conf.Database.Charset,
	)
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: conf.Database.TablePrefix,
			//SingularTable: true,

		},
		DisableForeignKeyConstraintWhenMigrating: true,      // 禁用自动创建外键约束
		Logger:                                   newLogger, // 使用自定义 Logger
	}); err != nil {
		gLog.Error("failed opening connection to err:", zap.Any("err", err))
		panic("failed to connect database")
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(conf.Database.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.Database.MaxOpenConns)
		return db
	}
}

// redis数据库连接

// NewRedis .
func NewRedis(c *config.Configuration, gLog *zap.Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host + ":" + c.Redis.Port,
		Password: c.Redis.Password, // no password set
		DB:       c.Redis.DB,       // use default DB
	})

	client.AddHook(redisotel.TracingHook{})
	if err := client.Ping(context.Background()).Err(); err != nil {
		gLog.Error("redis connect failed, err:", zap.Any("err", err))
		panic("failed to connect redis")
	}

	return client
}

// NewMongoDB mongo数据库连接
func NewMongoDB(conf *config.Configuration, gLog *zap.Logger) *mongo.Database {

	//dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("mongodb://%s:%s",
		conf.MongoDB.Host,
		strconv.Itoa(conf.MongoDB.Port),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// client option
	var clientOpt = options.Client().ApplyURI(dsn)

	zapLogger := newLogger(conf, conf.MongoDB.LogFilename)

	// log monitor
	var logMonitor = event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {

			msg := fmt.Sprintf("mongo reqId:%d start on db:%s cmd:%s sql:%+v", startedEvent.RequestID, startedEvent.DatabaseName,
				startedEvent.CommandName, startedEvent.Command)
			zapLogger.Info(msg)
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {
			msg := fmt.Sprintf("mongo reqId:%d exec cmd:%s success duration %d ns", succeededEvent.RequestID,
				succeededEvent.CommandName, succeededEvent.DurationNanos)
			zapLogger.Info(msg)
		},
		Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {
			msg := fmt.Sprintf("mongo reqId:%d exec cmd:%s failed duration %d ns", failedEvent.RequestID,
				failedEvent.CommandName, failedEvent.DurationNanos)
			zapLogger.Error(msg)
		},
	}

	// cmd monitor set
	clientOpt.SetMonitor(&logMonitor)

	//1.建立连接
	mongoClient, err := mongo.Connect(ctx, clientOpt)
	if nil != err {
		fmt.Printf("mongo connect err %v\n", err)
	} else {
		//fmt.Printf("mongo connect success~\n")
		zapLogger.Info("mongo connect success~")
		defer func() {
			if err = mongoClient.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	}

	//mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn).SetConnectTimeout(10*time.Second))
	database := mongoClient.Database(conf.MongoDB.Database)
	if err != nil {
		//fmt.Print(err)
		gLog.Error("mongo connect failed, err:", zap.Any("err", err))
		panic("failed to connect mongo")
	}

	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		gLog.Error("mongo connect failed, err:", zap.Any("err", err))
		panic("failed to connect mongo")
	}
	return database
}

type contextTxKey struct{}

// 事务确保一致性

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

// 新建数据库事务

// NewTransaction .
func NewTransaction(d *Data) service.Transaction {
	return d
}

// 初始化日志
func newLogger(conf *config.Configuration, filename string) *zap.Logger {
	var rootPath = path.RootPath()
	var level zapcore.Level    // zap 日志等级
	var zapOption []zap.Option // zap 配置项

	logFileDir := conf.Log.RootDir
	fmt.Println(filename)

	//logFileDir := conf.Log.RootDir
	if !filepath.IsAbs(logFileDir) {
		logFileDir = filepath.Join(rootPath, logFileDir)
	}

	if ok, _ := path.Exists(logFileDir); !ok {
		_ = os.Mkdir(conf.Log.RootDir, os.ModePerm)
	}

	switch conf.Log.Level {
	case "debug":
		level = zap.DebugLevel
		zapOption = append(zapOption, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		zapOption = append(zapOption, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(conf.App.Env + "." + l.String())
	}

	loggerWriter := &lumberjack.Logger{
		Filename:   filepath.Join(logFileDir, filename),
		MaxSize:    conf.Log.MaxSize,
		MaxBackups: conf.Log.MaxBackups,
		MaxAge:     conf.Log.MaxAge,
		Compress:   conf.Log.Compress,
	}

	zapLogger := zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(loggerWriter), level), zapOption...)
	return zapLogger
}
