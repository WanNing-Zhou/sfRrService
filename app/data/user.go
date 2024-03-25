package data

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/model"
	"github.com/jassue/gin-wire/app/service"
	"go.uber.org/zap"
)

type userRepo struct {
	data *Data
	log  *zap.Logger
}

func NewUserRepo(data *Data, log *zap.Logger) service.UserRepo {
	return &userRepo{
		data: data,
		log:  log,
	}
}

func (r *userRepo) FindByID(ctx context.Context, id uint64) (*domain.User, error) {
	var user model.User
	if err := r.data.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	//fmt.Println('')
	return user.ToDomain(), nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user model.User
	if err := r.data.db.Where(&domain.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return user.ToDomain(), nil
}

func (r *userRepo) FindByMobile(ctx context.Context, mobile string) (*domain.User, error) {
	var user model.User

	if err := r.data.db.Where(&domain.User{Mobile: mobile}).First(&user).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *userRepo) Create(ctx context.Context, u *domain.User) (*domain.User, error) {
	var user model.User

	id, err := r.data.sf.NextID()
	if err != nil {
		return nil, err
	}
	user.ID = id
	user.Name = u.Name
	user.Mobile = u.Mobile
	user.Password = u.Password
	user.Email = u.Email

	if err = r.data.DB(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *userRepo) Update(ctx context.Context, u *domain.User) (*domain.User, error) {
	var user model.User
	//id, err := r.data.sf.NextID()
	//if err != nil { return nil, err}
	//user.Password = u.Password
	user.Email = u.Email
	user.Mobile = u.Mobile
	user.Name = u.Name
	user.ID = u.ID
	user.Avatar = u.Avatar
	user.Introduction = u.Introduction

	// 忽略password, email 进行更新
	if err := r.data.DB(ctx).Omit("password", "email").Updates(&user).Error; err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

// ReSetInfo 重设置密码
func (r *userRepo) ReSetInfo(ctx context.Context, u *domain.User) (*domain.User, error) {
	var user model.User
	// :TODO to implement
	return user.ToDomain(), nil
}
