package data

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/model"
	"github.com/jassue/gin-wire/app/service"
	"go.uber.org/zap"
)

type compRepo struct {
	data *Data
	log  *zap.Logger
}

func NewCompRepo(data *Data, log *zap.Logger) service.CompRepo {
	return &compRepo{
		data: data,
		log:  log,
	}
}

// 根据id寻找comp

func (r *compRepo) FindByID(ctx context.Context, id uint64) (*domain.Comp, error) {
	var comp model.Comp
	if err := r.data.db.First(&comp, id).Error; err != nil {
		return nil, err
	}
	return comp.ToDomain(), nil
}

// 根据创建者id寻找comp

func (r *compRepo) FindByCreateId(ctx context.Context, creatId uint64) (*domain.Comp, error) {
	var comp model.Comp
	if err := r.data.db.First(&comp, creatId).First(&comp).Error; err != nil {
		return nil, err
	}
	return comp.ToDomain(), nil
}

// 创建comp

func (r *compRepo) Create(ctx context.Context, c *domain.Comp) (*domain.Comp, error) {
	var comp model.Comp

	id, err := r.data.sf.NextID()
	if err != nil {
		return nil, err
	}
	comp.ID = id
	comp.Name = c.Name
	comp.Info = c.Info
	comp.IsList = 0

	if err = r.data.DB(ctx).Create(&comp).Error; err != nil {
		return nil, err
	}

	return comp.ToDomain(), nil
}