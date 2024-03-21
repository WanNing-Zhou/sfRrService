package model

import (
	"github.com/jassue/gin-wire/app/domain"
	"time"
)

type Comp struct {
	ID       uint64    `gorm:"primaryKey"` // 主键
	Name     string    `gorm:"Size:30;not null;comment:组件名称"`
	Info     string    `gorm:"Size:100;not null;comment:组件信息"`
	CreatId  uint64    `gorm:"not null;comment:创建人"`
	CreateAt time.Time `gorm:"not null;comment:创建时间"`
	UpdateAt time.Time `gorm:"not null;comment:更新时间"`
	IsList   int       `gorm:"not null;comment:是否上架"`
}

// ToDomain 转换为领域模型
func (c *Comp) ToDomain() *domain.Comp {
	return &domain.Comp{
		ID:       c.ID,
		Name:     c.Name,
		Info:     c.Info,
		CreatId:  c.CreatId,
		CreateAt: c.CreateAt,
		UpdateAt: c.UpdateAt,
		IsList:   c.IsList,
	}
}
