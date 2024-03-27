package service

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
	cErr "github.com/jassue/gin-wire/app/pkg/error"
	"github.com/jassue/gin-wire/app/pkg/request"
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

func (s *CompService) Create(ctx context.Context, param *request.NewComp) (*domain.Comp, error) {

	u, err := s.cRepo.Create(ctx, &domain.Comp{
		Title:      param.Title,
		Info:       param.Info,
		Deploy:     param.Deploy,
		Types:      param.Types,
		PreviewUrl: param.PreviewUrl,
		Url:        param.Url,
		CreateId:   param.CreateId,
		Row:        param.Row,
		Column:     param.Column,
	})
	if err != nil {
		return nil, cErr.BadRequest("创建失败")
	}

	return u, nil
}
