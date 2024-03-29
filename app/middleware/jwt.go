package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jassue/gin-wire/app/domain"
	cErr "github.com/jassue/gin-wire/app/pkg/error"
	"github.com/jassue/gin-wire/app/pkg/response"
	"github.com/jassue/gin-wire/app/service"
	"github.com/jassue/gin-wire/config"
	"github.com/jassue/gin-wire/util/str"
	"strconv"
	"time"
)

type JWTAuth struct {
	conf *config.Configuration
	jwtS *service.JwtService
}

func NewJWTAuthM(
	conf *config.Configuration,
	jwtS *service.JwtService,
) *JWTAuth {
	return &JWTAuth{
		conf: conf,
		jwtS: jwtS,
	}
}

func (m *JWTAuth) Handler(guardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		tokenStr := str.SplitToken(bearToken)
		//fmt.Println(tokenStr)
		if tokenStr == "" {
			response.FailByErr(c, cErr.Unauthorized("missing Authorization header"))
			return
		}

		token, err := jwt.ParseWithClaims(tokenStr, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.conf.Jwt.Secret), nil
		})

		if err != nil || m.jwtS.IsInBlacklist(c, tokenStr) {
			response.FailByErr(c, cErr.Unauthorized("登录授权已失效"))
			return
		}

		//fmt.Println(guardName)
		claims := token.Claims.(*domain.CustomClaims)
		//fmt.Println(token.Claims.(*domain.CustomClaims).Key)

		//if claims.Issuer != guardName {
		if claims.Key != guardName {
			fmt.Println("groundName err")
			response.FailByErr(c, cErr.Unauthorized("登录授权已失效"))
			return
		}

		// token 续签
		if int64(claims.ExpiresAt.Sub(time.Now()).Seconds()) < m.conf.Jwt.RefreshGracePeriod {
			tokenData, err := m.jwtS.RefreshToken(c, guardName, token)
			if err == nil {
				c.Header("new-token", tokenData.AccessToken)
				c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
			}
		}

		c.Set("token", token)
		c.Set("id", claims.ID)
		c.Set("auth", claims.Auth)
	}
}

func (m *JWTAuth) AuthDevHandle(guardName string) gin.HandlerFunc {

	return func(c *gin.Context) {
		auth, err := strconv.Atoi(c.Keys["auth"].(string))
		if auth > 2 || err != nil {
			response.FailByErr(c, cErr.Forbidden("权限不足"))
			return
		}

	}

}

// AuthSuperHandle  验证管理权限
func (m *JWTAuth) AuthSuperHandle(groundName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth, err := strconv.Atoi(c.Keys["auth"].(string))
		if auth > 1 || err != nil {
			response.FailByErr(c, cErr.Forbidden("权限不足"))
			return
		}

	}
}
