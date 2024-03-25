package model

import (
	"github.com/jassue/gin-wire/app/domain"
)

type User struct {
	ID           uint64 `gorm:"primaryKey"`
	Name         string `gorm:"size:30;not null;comment:用户名称"`
	Mobile       string `gorm:"size:24;not null;index;comment:用户手机号"`
	Password     string `gorm:"not null;default:'';comment:用户密码"`
	Email        string `gorm:"size:30;not null;comment:用户邮箱"`
	Avatar       string `gorm:"size:255;not null;comment:头像"`
	Introduction string `gorm:"size:255;comment:个人介绍"`
	Auth         int    `gorm:"not null;default:0;comment:权限"`
	Timestamps
	//SoftDeletes
}

func (m *User) ToDomain() *domain.User {
	return &domain.User{
		ID:           m.ID,
		Name:         m.Name,
		Mobile:       m.Mobile,
		Email:        m.Email,
		Password:     m.Password,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		Avatar:       m.Avatar,
		Introduction: m.Introduction,
	}
}
