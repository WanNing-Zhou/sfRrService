package data

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/model"
	"github.com/jassue/gin-wire/app/service"
	"go.uber.org/zap"
)

type cAMsgRepo struct {
	data *Data
	log  *zap.Logger
}

func NewCAMsgRepo(data *Data, log *zap.Logger) service.CAMsgRepo {
	return &cAMsgRepo{
		data: data,
		log:  log,
	}
}

func (r *cAMsgRepo) Create(ctx context.Context, msg *domain.CompAuditMsg) (*domain.CompAuditMsg, error) {
	var resMsg model.CompAuditMsg

	id, err := r.data.sf.NextID()
	if err != nil {
		return nil, err
	}
	resMsg.ID = id
	resMsg.CreateId = msg.CreateId
	resMsg.IsList = msg.IsList
	resMsg.CompId = msg.CompId
	resMsg.Msg = msg.Msg

	if err = r.data.DB(ctx).Create(&resMsg).Error; err != nil {
		return nil, err
	}

	return resMsg.ToDomain(), nil
}
