package service

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
	cErr "github.com/jassue/gin-wire/app/pkg/error"
	"github.com/jassue/gin-wire/app/pkg/request"
)

type PageRepo interface {
	FindByID(context.Context, uint64) (*domain.Comp, error)
	FindByCreateId(context.Context, uint64) (*domain.Comp, error) // 根据创建人ID查找
	Create(context.Context, *domain.Page) (*domain.Page, error)   // 创建
	FindCompsByQuery(context.Context, *request.CompList, bool) ([]domain.Comp, int64, error)
	UpdateComp(ctx context.Context, comp *domain.Comp) (*domain.Comp, error)
	UpdateIsList(ctx context.Context, c *domain.Comp) (*domain.Comp, error)
}

type PageService struct {
	pRepo PageRepo
	tm    Transaction
	cAM   CAMsgRepo
}

func NewPageService(pRepo PageRepo, tm Transaction, cAMRepo CAMsgRepo) *PageService {
	return &PageService{
		pRepo: pRepo,
		tm:    tm,
		cAM:   cAMRepo,
	}
}

func (s *PageService) Create(ctx context.Context, param *request.NewPage) (*domain.Page, error) {

	u, err := s.pRepo.Create(ctx, &domain.Page{
		Title:    param.Title,
		Info:     param.Info,
		CreateId: param.CreateId,
		Data:     param.Data,
	})
	if err != nil {
		return nil, cErr.BadRequest("创建失败")
	}

	return u, nil
}
