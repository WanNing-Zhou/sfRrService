package domain

import (
	"strconv"
	"time"
)

type Comp struct {
	ID         uint64    `json:"id,string"` // 主键
	Title      string    `json:"title"`
	Info       string    `json:"info"`
	CreateId   uint64    `json:"create_id,string"`
	CreatedAt  time.Time `json:"create_at"`
	UpdatedAt  time.Time `json:"UpdateAt"`
	IsList     int       `json:"IsList"`
	Deploy     string    `json:"deploy"`
	Types      int       `json:"types"`
	PreviewUrl string    `json:"preview_url"`
	Url        string    `json:"url"`
	Row        int       `json:"row"`
	Column     int       `json:"column"`
}

// 获取组件id

func (c *Comp) GetCid() string {
	return strconv.Itoa(int(c.ID))
}
