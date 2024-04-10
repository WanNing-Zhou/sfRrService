package data

import (
	"context"
	"fmt"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/service"
	"github.com/jassue/gin-wire/util/paginate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"time"
)

type pageRepo struct {
	data *Data
	log  *zap.Logger
}

func NewPageRepo(data *Data, log *zap.Logger) service.PageRepo {
	return &pageRepo{
		data: data,
		log:  log,
	}
}

func (p pageRepo) FindByID(ctx context.Context, id string) (*domain.Page, error) {
	//TODO implement me
	collections := p.data.mdb.Collection("page")
	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Println("objID", objID)
	var res *domain.Page
	err := collections.FindOne(ctx, bson.M{"_id": objID}).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p pageRepo) FindByCreateId(ctx context.Context, u uint64) (*domain.Comp, error) {
	//TODO implement me
	panic("implement me")
}

// Create 创建
func (p pageRepo) Create(ctx context.Context, comp *domain.Page) error {

	//ash := bson.M{"name": "kkx"}
	comp.CreatedAt = time.Now()
	comp.UpdatedAt = time.Now()
	collections := p.data.mdb.Collection("page")
	_, err := collections.InsertOne(context.TODO(), comp)

	if err != nil {
		return err
	}
	return nil
}

// FindPagesByQuery 查询
func (p pageRepo) FindPagesByQuery(ctx context.Context, params *request.GetPages, b bool) ([]domain.Page, int64, error) {

	findOptions := paginate.MPaginate(&request.PageDto{Page: params.Page, PageSize: params.PageSize})
	collections := p.data.mdb.Collection("page")
	filter := bson.D{}
	if params.Title != "" {
		filter = append(filter, bson.E{
			Key: "title",
			//i 表示不区分大小写
			Value: bson.M{"$regex": primitive.Regex{Pattern: ".*" + params.Title + ".*", Options: "i"}},
		})
	}
	if params.ID != "" {
		filter = append(filter, bson.E{
			Key:   "_id",
			Value: params.ID,
		})
	}
	if params.CreateId != 0 {
		filter = append(filter, bson.E{
			Key:   "create_id",
			Value: params.CreateId,
		})
	}
	count, err2 := collections.CountDocuments(ctx, &filter)
	if err2 != nil {
		return nil, 0, err2
	}

	cur, err := collections.Find(ctx, &filter, findOptions)

	if err != nil {
		return nil, 0, err
	}

	var res []domain.Page
	//var rss []bson.M
	err = cur.All(ctx, &res)

	if err != nil {
		return nil, 0, err
	}

	//fmt.Println(res)

	return res, count, nil
}

func (p pageRepo) UpdateComp(ctx context.Context, comp *domain.Comp) (*domain.Comp, error) {
	//TODO implement me
	panic("implement me")
}

func (p pageRepo) UpdateIsList(ctx context.Context, c *domain.Comp) (*domain.Comp, error) {
	//TODO implement me
	panic("implement me")
}
