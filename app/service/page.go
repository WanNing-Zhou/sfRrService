package service

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
	cErr "github.com/jassue/gin-wire/app/pkg/error"
	"github.com/jassue/gin-wire/app/pkg/request"
)

type PageRepo interface {
	FindByID(context.Context, string) (*domain.Page, error)
	FindByCreateId(context.Context, uint64) (*domain.Comp, error) // 根据创建人ID查找
	Create(context.Context, *domain.Page) error                   // 创建
	FindPagesByQuery(context.Context, *request.GetPages, bool) ([]domain.Page, int64, error)
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

func (s *PageService) Create(ctx context.Context, param *request.NewPage) error {

	err := s.pRepo.Create(ctx, &domain.Page{
		Title:    param.Title,
		Info:     param.Info,
		CreateId: param.CreateId,
		Data:     param.Data,
	})
	if err != nil {
		return cErr.BadRequest("创建失败")
	}

	return nil
}

func (s *PageService) GetPages(ctx context.Context, param *request.GetPages) ([]domain.Page, int64, error) {

	query, i, err := s.pRepo.FindPagesByQuery(ctx, param, true)
	if err != nil {
		return nil, 0, cErr.BadRequest("查询失败")
	}

	return query, i, nil
}

func (s *PageService) FindById(ctx context.Context, id string) (*domain.Page, error) {
	page, err := s.pRepo.FindByID(ctx, id)

	if err != nil {
		return nil, cErr.BadRequest("查询失败")
	}

	return page, nil
}
