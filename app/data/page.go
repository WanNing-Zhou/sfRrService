package data

import (
	"context"
	"fmt"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
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

func (p pageRepo) FindByID(ctx context.Context, u uint64) (*domain.Comp, error) {
	//TODO implement me
	panic("implement me")
}

func (p pageRepo) FindByCreateId(ctx context.Context, u uint64) (*domain.Comp, error) {
	//TODO implement me
	panic("implement me")
}

func (p pageRepo) Create(ctx context.Context, comp *domain.Page) (*domain.Page, error) {

	ash := bson.M{"name": "kkx"}
	collections := p.data.mdb.Collection("myCollection")
	//findAll, _ := collections.Find(context.TODO(), nil)
	one, err := collections.InsertOne(context.TODO(), ash)
	if err != nil {
		return nil, err
	}

	fmt.Println(one)
	fmt.Println("执行了")
	//collections.Collection("trainers")
	//TODO implement me
	panic("implement me")
}

func (p pageRepo) FindCompsByQuery(ctx context.Context, list *request.CompList, b bool) ([]domain.Comp, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (p pageRepo) UpdateComp(ctx context.Context, comp *domain.Comp) (*domain.Comp, error) {
	//TODO implement me
	panic("implement me")
}

func (p pageRepo) UpdateIsList(ctx context.Context, c *domain.Comp) (*domain.Comp, error) {
	//TODO implement me
	panic("implement me")
}
