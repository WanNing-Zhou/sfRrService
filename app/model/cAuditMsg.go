package model

import "github.com/jassue/gin-wire/app/domain"

type CompAuditMsg struct {
	ID       uint64 `gorm:"primaryKey"` // 主键
	CompId   uint64 `gorm:"comment:组件ID"`
	Msg      string `gorm:"Size:255;comment:消息"`
	CreateId uint64 `gorm:"comment:创建人"`
	IsList   int    `gorm:"组组件状态"`
	Timestamps
}

// ToDomain 转换为领域模型
func (c *CompAuditMsg) ToDomain() *domain.CompAuditMsg {
	return &domain.CompAuditMsg{
		ID:        c.ID,
		CompId:    c.CompId,
		CreateId:  c.CreateId,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		IsList:    c.IsList,
		Msg:       c.Msg,
	}
}
