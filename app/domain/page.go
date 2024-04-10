package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Page struct {
	CreateId  uint64             `json:"create_id,string" bson:"create_id,omitempty"` // 创建人
	ID        primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`              // 页面id
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`                // 创建时间
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`                // 更新时间
	Data      string             `json:"data,string" bson:"data,string"`              // 页面数据
	Info      string             `json:"info" bson:"info"`                            // 页面信息
	Title     string             `json:"title" bson:"title"`                          // 页面标题
}
