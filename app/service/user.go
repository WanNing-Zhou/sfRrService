package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jassue/gin-wire/app/domain"
	cErr "github.com/jassue/gin-wire/app/pkg/error"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/util/hash"
	"strconv"
)

type UserRepo interface {
	FindByID(context.Context, uint64) (*domain.User, error)
	FindByMobile(context.Context, string) (*domain.User, error)
	Create(context.Context, *domain.User) (*domain.User, error)
	FindByEmail(context.Context, string) (*domain.User, error)               // 根据Email寻找用户
	Update(context.Context, *domain.User) (*domain.User, error)              // 更新用户信息
	UpdatePassword(context.Context, *request.Password) (*domain.User, error) // 设置密码
	FindByQuery(ctx context.Context, users *request.GetUsers) ([]domain.User, int64, error)
}

type UserService struct {
	uRepo UserRepo
	tm    Transaction
}

// NewUserService .
func NewUserService(uRepo UserRepo, tm Transaction) *UserService {
	return &UserService{uRepo: uRepo, tm: tm}
}

// Register 注册
func (s *UserService) Register(ctx *gin.Context, param *request.Register) (*domain.User, error) {
	//user, _ := s.uRepo.FindByMobile(ctx, param.Mobile)
	user, _ := s.uRepo.FindByEmail(ctx, param.Email)
	if user != nil {
		return nil, cErr.BadRequest("该邮箱已被注册")
	}

	// 创建用户
	u, err := s.uRepo.Create(ctx, &domain.User{
		Name: param.Name,
		//Mobile:   param.Mobile,
		Password: hash.BcryptMake([]byte(param.Password)),
		Email:    param.Email,
	})
	if err != nil {
		return nil, cErr.BadRequest("注册用户失败")
	}

	return u, nil
}

// SetInfo 设置用户信息
func (s *UserService) SetInfo(ctx *gin.Context, param *request.Info) (*domain.User, error) {

	u, err := s.uRepo.Update(ctx, &domain.User{
		Name:         param.Name,
		Mobile:       param.Mobile,
		ID:           param.ID,
		Email:        param.Email,
		Avatar:       param.Avatar,
		Introduction: param.Introduction,
	})
	if err != nil {
		return nil, cErr.BadRequest("设置用户信息失败")
	}

	return u, nil
}

func (s *UserService) SetPassword(ctx *gin.Context, param *request.Password) (*domain.User, error) {
	u, err := s.uRepo.FindByID(ctx, param.ID)

	if err != nil || !hash.BcryptMakeCheck([]byte(param.OldPassword), u.Password) {
		return nil, cErr.BadRequest("旧密码错误")
	}

	nu, err := s.uRepo.UpdatePassword(ctx, &request.Password{
		ID:          param.ID,
		NewPassword: hash.BcryptMake([]byte(param.NewPassword)),
	})
	if err != nil {
		return nil, cErr.BadRequest("设置用户密码失败")
	}
	return nu, nil
}

// Login 登录
func (s *UserService) Login(ctx *gin.Context, email, password string) (*domain.User, error) {
	//u, err := s.uRepo.FindByMobile(ctx, mobile)
	// 根据邮箱查找账户
	u, err := s.uRepo.FindByEmail(ctx, email)
	if err != nil || !hash.BcryptMakeCheck([]byte(password), u.Password) {
		return nil, cErr.BadRequest("用户名不存在或密码错误")
	}

	return u, nil
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(ctx *gin.Context, idStr string) (*domain.User, error) {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, cErr.NotFound("数据ID错误")
	}
	u, err := s.uRepo.FindByID(ctx, id)
	if err != nil {
		return nil, cErr.NotFound("数据不存在", cErr.USER_NOT_FOUND)
	}
	//fmt.Print(u)
	return u, nil
}

// SLogin 管理员登陆
func (s *UserService) SLogin(ctx *gin.Context, email, password string) (*domain.User, error) {
	//u, err := s.uRepo.FindByMobile(ctx, mobile)
	// 根据邮箱查找账户
	u, err := s.uRepo.FindByEmail(ctx, email)
	if err != nil || !hash.BcryptMakeCheck([]byte(password), u.Password) {
		return nil, cErr.BadRequest("用户名不存在或密码错误")
	}
	if u.Auth > 0 {
		return nil, cErr.BadRequest("该用户无管理员权限")
	}
	return u, nil
}

func (s *UserService) GetUsers(ctx *gin.Context, params *request.GetUsers) ([]domain.User, int64, error) {
	// 根据分页信息查询用户
	users, total, err := s.uRepo.FindByQuery(ctx, params)
	if err != nil {
		return nil, 0, cErr.BadRequest("查询失败")
	}

	return users, total, nil
}
