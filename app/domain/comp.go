package domain

import (
	"strconv"
	"time"
)

type Comp struct {
	ID       uint64    `json:"id"` // 主键
	Name     string    `json:"name"`
	Info     string    `json:"info"`
	CreatId  uint64    `json:"creat_id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"UpdateAt"`
	IsList   int       `json:"IsList"`
}

// 获取组件id

func (c *Comp) GetCid() string {
	return strconv.Itoa(int(c.ID))
}
