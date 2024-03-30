package domain

import "time"

type CompAuditMsg struct {
	ID        uint64    `json:"id,string"` // 主键
	CompId    uint64    `json:"comp_id,string"`
	Msg       string    `json:"msg"`
	CreateId  uint64    `json:"create_id,string"`
	IsList    int       `json:"is_list"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
