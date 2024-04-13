package data

import (
	"context"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/model"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/service"
	"github.com/jassue/gin-wire/util/paginate"
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

func (r *userRepo) DeleteByID(ctx context.Context, id uint64) error {
	var user model.User
	user.ID = id
	if err := r.data.db.Model(user).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// FindByQuery 根据条件查询用户列表
func (r *userRepo) FindByQuery(ctx context.Context, param *request.GetUsers) ([]domain.User, int64, error) {
	var user model.User

	query := r.data.db.Model(&user)

	// 对名字进行模糊查询
	if param.Name != "" {
		query = query.Where("name LIKE ?", "%"+param.Name+"%")
	}

	if param.ID != 0 {
		query = query.Where("id = ?", param.ID)
	}
	if param.Email != "" {
		query = query.Where("email = ?", param.Email)
	}

	// 排序
	query = query.Order("updated_at desc")

	pageReq := &request.PageDto{
		Page:     param.Page,
		PageSize: param.PageSize,
	}

	var userList []model.User
	var total int64
	// 获取获取过滤后数据的总条数
	query.Count(&total)

	err := query.Scopes(paginate.Paginate(pageReq)).Find(&userList).Error
	if err != nil {
		return nil, 0, err
	}

	var users []domain.User

	for _, v := range userList {
		v.Password = ""
		users = append(users, *v.ToDomain())
	}

	return users, total, nil
}

func (r *userRepo) UpdatePassword(ctx context.Context, params *request.Password) (*domain.User, error) {
	var user model.User

	user.ID = params.ID
	user.Password = params.NewPassword

	if err := r.data.DB(ctx).Updates(&user).Error; err != nil {
		return nil, err
	}

	//fmt.Print("user2", user)

	return &domain.User{
		ID:   user.ID,
		Name: user.Name,
		Auth: user.Auth,
	}, nil
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
	user.Auth = u.Auth

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
