package domain

import (
	"strconv"
	"time"
)

type User struct {
	ID           uint64    `json:"id,string"`          // 用户ID
	Name         string    `json:"name"`               // 用户名
	Email        string    `json:"email"`              // 邮箱
	Mobile       string    `json:"mobile"`             // 手机号
	Password     string    `json:"password,omitempty"` // 密码
	CreatedAt    time.Time `json:"created_at"`         // 创建时间
	UpdatedAt    time.Time `json:"updated_at"`         // 更新时间
	Avatar       string    `json:"avatar"`             // 头像
	NickName     string    `json:"nick_name"`          // 昵称
	Introduction string    `json:"introduction"`       // 个人简介
	Auth         int       `'json:"auth"`              // 权限
}

// 获取用户UID

func (u *User) GetUid() string {
	return strconv.Itoa(int(u.ID))
}

// GetAuth 获取用户权限

func (u *User) GetAuth() string {
	return strconv.Itoa(u.Auth)
}
