package domain

import (
	"strconv"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`         // 用户ID
	Name      string    `json:"name"`       // 用户名
	Email     string    `json:"email"`      // 邮箱
	Mobile    string    `json:"mobile"`     // 手机号
	Password  string    `json:"password"`   // 密码
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	Avatar    string    `json:"avatar"`     // 头像
	NickName  string    `json:"nick_name"`  // 昵称
}

// 获取用户UID

func (u *User) GetUid() string {
	return strconv.Itoa(int(u.ID))
}
