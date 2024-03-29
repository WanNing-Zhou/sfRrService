package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jassue/gin-wire/app/domain"
	"github.com/jassue/gin-wire/app/handler/app"
	"github.com/jassue/gin-wire/app/handler/common"
	"github.com/jassue/gin-wire/app/middleware"
)

// 设置api留有

func setApiGroupRoutes(
	router *gin.Engine,
	jwtAuthM *middleware.JWTAuth,
	authH *app.AuthHandler,
	commonH *common.UploadHandler,
	compH *app.CompHandler,
) *gin.RouterGroup {
	group := router.Group("/api")
	group.POST("/auth/register", authH.Register)
	group.POST("/auth/login", authH.Login)
	// 权限校验
	authGroup := group.Group("").Use(jwtAuthM.Handler(domain.AppGuardName))
	{
		authGroup.GET("/auth/info", authH.Info)
		authGroup.POST("/auth/logout", authH.Logout)
		authGroup.POST("/image_upload", commonH.ImageUpload)
		authGroup.POST("/auth/info", authH.SetInfo)
		authGroup.POST("/auth/password", authH.SetPassword)
	}

	// 需要开发者以上权限
	compGroup := group.Group("/comp").Use(jwtAuthM.Handler(domain.AppGuardName)).Use(jwtAuthM.AuthDevHandle(domain.AppGuardName))
	{
		compGroup.POST("/create", compH.NewComp)
		compGroup.GET("/list", compH.GetCompList)
		compGroup.GET("/info", compH.GetCompInfo)
		compGroup.POST("/update", compH.UpdateCompInfo)
	}

	// 管理员登陆
	group.POST("/sAuth/login", authH.SLogin)

	// 需要管理员权限
	sAuthGroup := group.Group("/sAuth").Use(jwtAuthM.Handler(domain.AppGuardName)).Use(jwtAuthM.AuthSuperHandle(domain.AppGuardName))
	{
		// 获取用户列表
		sAuthGroup.GET("/user/list")
		// 创建管理员账号
		sAuthGroup.POST("/sUser/create")
		// 组件审核
		sAuthGroup.POST("/comp/audit")
		// 组件删除
		sAuthGroup.POST("/comp/delete")
	}

	return group
}
