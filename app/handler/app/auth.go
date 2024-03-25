package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/pkg/request"
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

}
