package model

import (
	"github.com/jassue/gin-wire/app/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Page struct {
	CreateId  uint64             `json:"create_id,string" bson:"create_id"` // 创建人
	ID        primitive.ObjectID `json:"id" bson:"_id"`                     // 页面id
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`      // 创建时间
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`      // 更新时间
	Data      interface{}        `json:"data" bson:"data"`                  // 页面数据
	Info      string             `json:"info" bson:"info"`                  // 页面信息
	Title     string             `json:"title" bson:"title"`                // 页面标题
}

func (p *Page) ToDomain() *domain.Page {
	return &domain.Page{
		ID:        p.ID,
		Title:     p.Title,
		Info:      p.Info,
		Data:      p.Data,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		CreateId:  p.CreateId,
	}
}
