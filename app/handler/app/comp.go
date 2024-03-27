package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/pkg/response"
	"github.com/jassue/gin-wire/app/service"
	"go.uber.org/zap"
	"strconv"
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

// NewComp 创建Comp
func (h *CompHandler) NewComp(c *gin.Context) {

	var form request.NewComp

	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	userId, _ := strconv.ParseUint(c.Keys["id"].(string), 10, 64)

	form.CreateId = userId

	//fmt.Println("kkx")

	//fmt.Println(form)
	comp, err := h.compS.Create(c, &form)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, comp)
}
