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
	FindCompsByQuery(context.Context, *request.CompList) ([]domain.Comp, int64, error)
	UpdateComp(ctx context.Context, comp *domain.Comp) (*domain.Comp, error)
	UpdateIsList(ctx context.Context, c *domain.Comp) (*domain.Comp, error)
}

type CompService struct {
	cRepo CompRepo
	tm    Transaction
	cAM   CAMsgRepo
}

func NewCompService(cRepo CompRepo, tm Transaction, cAMRepo CAMsgRepo) *CompService {
	return &CompService{
		cRepo: cRepo,
		tm:    tm,
		cAM:   cAMRepo,
	}
}

func (s *CompService) Create(ctx context.Context, param *request.NewComp) (*domain.Comp, error) {

	u, err := s.cRepo.Create(ctx, &domain.Comp{
		Title:  param.Title,
		Info:   param.Info,
		Deploy: param.Deploy,
		//Types:      param.Types,
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

func (s *CompService) GetComps(ctx context.Context, param *request.CompList) ([]domain.Comp, int64, error) {
	compList, total, err := s.cRepo.FindCompsByQuery(ctx, param)
	if err != nil {
		return nil, 0, cErr.BadRequest("查询失败")
	}

	return compList, total, nil
}

func (s *CompService) GetCompById(ctx context.Context, id uint64) (*domain.Comp, error) {
	comp, err := s.cRepo.FindByID(ctx, id)
	if err != nil {
		return nil, cErr.BadRequest("查询失败")
	}
	return comp, nil
}

func (s *CompService) UpdateComp(ctx context.Context, param *request.UpdateCompInfo) (*domain.Comp, error) {
	var comp domain.Comp
	comp.ID = param.ID
	comp.Row = param.Row
	comp.Column = param.Column
	comp.Url = param.Url
	comp.PreviewUrl = param.PreviewUrl
	comp.Deploy = param.Deploy
	comp.Info = param.Info
	comp.Title = param.Title
	resC, err := s.cRepo.UpdateComp(ctx, &comp)
	if err != nil {
		return nil, err
	}

	return resC, nil
}

func (s *CompService) AuditComp(ctx context.Context, param *request.AuditComp) (*domain.CompAuditMsg, error) {
	var comp domain.Comp
	comp.ID = param.ID
	comp.IsList = param.IsList
	_, err := s.cRepo.UpdateIsList(ctx, &comp)
	if err != nil {
		return nil, err
	}

	msg, err := s.cAM.Create(ctx, &domain.CompAuditMsg{
		CompId:   param.ID,
		CreateId: param.CreateId,
		Msg:      param.Msg,
		IsList:   param.IsList,
	})

	if err != nil {
		return nil, err
	}

	return msg, nil
}
