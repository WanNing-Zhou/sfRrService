package data

import (
	"context"
	"github.com/jassue/gin-wire/app/service"
	"github.com/jassue/gin-wire/util/hash"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type jwtRepo struct {
	data *Data
	log  *zap.Logger
}

// NewJwtRepo 新建jwt实例
func NewJwtRepo(data *Data, log *zap.Logger) service.JwtRepo {
	return &jwtRepo{
		data: data,
		log:  log,
	}
}

// GetBlackListKey 获取黑名单key
func (r *jwtRepo) getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + hash.MD5([]byte(tokenStr))
}

// JoinBlackList 加入黑名单
func (r *jwtRepo) JoinBlackList(ctx context.Context, tokenStr string, joinUnix int64, expires time.Duration) error {
	return r.data.rdb.SetNX(ctx, r.getBlackListKey(tokenStr), joinUnix, expires).Err()
}

// GetBlackJoinUnix 获取加入黑名单的时间
func (r *jwtRepo) GetBlackJoinUnix(ctx context.Context, tokenStr string) (int64, error) {
	joinUnixStr, err := r.data.rdb.Get(ctx, r.getBlackListKey(tokenStr)).Result()
	if err != nil {
		return 0, err
	}

	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return 0, err
	}

	return joinUnix, nil
}
