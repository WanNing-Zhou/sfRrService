package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/pkg/resp"
	"github.com/jassue/gin-wire/app/pkg/response"
	"github.com/jassue/gin-wire/app/service"
	"go.uber.org/zap"
)

type AuthHandler struct {
	log   *zap.Logger
	jwtS  *service.JwtService
	userS *service.UserService
}

func NewAuthHandler(log *zap.Logger, jwtS *service.JwtService, userS *service.UserService) *AuthHandler {
	return &AuthHandler{log: log, jwtS: jwtS, userS: userS}
}

// 用户注册

func (h *AuthHandler) Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	// 用户注册时直接赋予开发者权限
	form.Auth = 1

	u, err := h.userS.Register(c, &form)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	tokenData, _, err := h.jwtS.CreateToken(domain.AppGuardName, u)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, tokenData)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	user, err := h.userS.Login(c, form.Email, form.Password)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	tokenData, _, err := h.jwtS.CreateToken(domain.AppGuardName, user)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, tokenData)
}

func (h *AuthHandler) Info(c *gin.Context) {
	user, err := h.userS.GetUserInfo(c, c.Keys["id"].(string))
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	//delete(user, 'Passwrod')
	user.Password = ""
	response.Success(c, user)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	err := h.jwtS.JoinBlackList(c, c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, nil)
}

func (h AuthHandler) SetInfo(c *gin.Context) {
	var form request.Info
	// 表单校验
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	u, err := h.userS.SetInfo(c, &form)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, u)
}

func (h AuthHandler) SetPassword(c *gin.Context) {
	var form request.Password
	// 表单校验
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	u, err := h.userS.SetPassword(c, &form)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	// 密码设置成功后将原有的token加入黑名单
	blErr := h.jwtS.JoinBlackList(c, c.Keys["token"].(*jwt.Token))
	if blErr != nil {
		response.FailByErr(c, err)
		return
	}

	// 设置新token
	tokenData, _, jwtErr := h.jwtS.CreateToken(domain.AppGuardName, u)

	if jwtErr != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, tokenData)

	//response.Success(c, u)
}

// SLogin 管理员登陆
func (h AuthHandler) SLogin(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	user, err := h.userS.SLogin(c, form.Email, form.Password)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	tokenData, _, err := h.jwtS.CreateToken(domain.AppGuardName, user)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, tokenData)
}

func (h AuthHandler) GetUserList(c *gin.Context) {
	var form request.GetUsers
	if err := c.ShouldBindQuery(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	users, total, err := h.userS.GetUsers(c, &form)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, &resp.RespList{
		List:  users,
		Total: total,
	})
}

func (h AuthHandler) GetUserInfo(c *gin.Context) {
	var form request.GetUserInfo
	if err := c.ShouldBindQuery(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	user, err := h.userS.GetUserInfo(c, form.ID)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	response.Success(c, user)
}

func (h AuthHandler) CreatSUser(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	form.Auth = 0
	_, err := h.userS.Register(c, &form)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, resp.SuccessData{
		Message: "创建成功",
	})

}

func (h AuthHandler) DeleteUser(c *gin.Context) {
	var form request.DeleteUser
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	h.userS.DeleteUser(c, form.ID)

}
