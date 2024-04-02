package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Page struct {
	CreateId  uint64             `json:"create_id,string"` // 创建人
	ID        primitive.ObjectID `json:"id,string"`        // 页面id
	CreatedAt time.Time          `json:"created_at"`       // 创建时间
	UpdatedAt time.Time          `json:"updated_at"`       // 更新时间
	Data      interface{}        `json:"data"`             // 页面数据
	Info      string             `json:"info"`             // 页面信息
	Title     string             `json:"title"`            // 页面标题
}
