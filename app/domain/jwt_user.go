package domain

import "github.com/golang-jwt/jwt/v4"

const (
	AppGuardName = "app"
)

type JwtUser interface {
	GetUid() string
	GetAuth() string
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	Key  string `json:"key,omitempty"`
	Auth string `json:"auth,omitempty"`
	jwt.RegisteredClaims
}

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
