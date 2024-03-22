package service

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
)

type CompRepo interface {
	FindByID(context.Context, uint64) (*domain.Comp, error)
	FindByCreateId(context.Context, uint64) (*domain.Comp, error) // 根据创建人ID查找
	Create(context.Context, *domain.Comp) (*domain.Comp, error)   // 创建
}

type CompService struct {
	cRepo CompRepo
	tm    Transaction
}

func NewCompService(cRepo CompRepo, tm Transaction) *CompService {
	return &CompService{
		cRepo: cRepo,
		tm:    tm,
	}
}
