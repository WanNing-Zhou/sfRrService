package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/pkg/response"
	"github.com/jassue/gin-wire/app/service"
	"go.uber.org/zap"
)

type PageHandler struct {
	log    *zap.Logger
	jwtS   *service.JwtService
	pageS  *service.PageService
	cAMsgS *service.CAMsgService
}

// NewPageHandler 创建CompHandler实例
func NewPageHandler(log *zap.Logger, jwtS *service.JwtService, pageS *service.PageService) *PageHandler {
	return &PageHandler{log: log, jwtS: jwtS, pageS: pageS}
}

// NewPage  创建Comp
func (h *PageHandler) NewPage(c *gin.Context) {

	var form request.NewPage

	//fmt.Println("kkx")

	//fmt.Println(form)
	comp, err := h.pageS.Create(c, &form)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, comp)
}
