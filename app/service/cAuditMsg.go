package service

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
)

type CAMsgRepo interface {
	Create(ctx context.Context, msg *domain.CompAuditMsg) (*domain.CompAuditMsg, error)
}

type CAMsgService struct {
	cAMRepo CAMsgRepo
}

func NewCAMsgService(cAMRepo CAMsgRepo) *CAMsgService {
	return &CAMsgService{
		cAMRepo: cAMRepo,
	}
}

func (s *CAMsgService) NewMsg(ctx context.Context, msg *domain.CompAuditMsg) (*domain.CompAuditMsg, error) {
	return s.cAMRepo.Create(ctx, msg)
}
