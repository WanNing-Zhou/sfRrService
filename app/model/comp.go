package model

import (
	"github.com/jassue/gin-wire/app/domain"
)

type Comp struct {
	ID         uint64 `gorm:"primaryKey"` // 主键
	Title      string `gorm:"Size:30;not null;comment:组件名称"`
	Info       string `gorm:"Size:100;not null;comment:组件信息"`
	CreateId   uint64 `gorm:"not null;comment:创建人"`
	IsList     int    `gorm:"comment:是否上架"`
	Deploy     string `gorm:"Size:100;comment:部署方式"`
	Url        string `gorm:"Size:100;comment:访问地址"`
	PreviewUrl string `gorm:"Size:100;comment:预览地址"`
	Types      int    `gorm:"comment:组件类型"`
	Timestamps
}

// ToDomain 转换为领域模型
func (c *Comp) ToDomain() *domain.Comp {
	return &domain.Comp{
		ID:         c.ID,
		Title:      c.Title,
		Info:       c.Info,
		CreateId:   c.CreateId,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		IsList:     c.IsList,
		Deploy:     c.Deploy,
		Url:        c.Url,
		PreviewUrl: c.PreviewUrl,
		Types:      c.Types,
	}
}
