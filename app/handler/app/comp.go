package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/pkg/response"
	"github.com/jassue/gin-wire/app/service"
	"go.uber.org/zap"
)

type CompHandler struct {
	log   *zap.Logger
	jwtS  *service.JwtService
	compS *service.CompService
}

// NewCompHandler  创建CompHandler实例
func NewCompHandler(log *zap.Logger, jwtS *service.JwtService, compS *service.CompService) *CompHandler {
	return &CompHandler{log: log, jwtS: jwtS, compS: compS}
}

func (h *CompHandler) NewComp(c *gin.Context) {

	var form request.NewComp
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	fmt.Println("kkx")

	fmt.Println(form)
	comp, err := h.compS.Create(c, &form)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, comp)
}
