package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/pkg/resp"
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

func (h *CompHandler) GetCompList(c *gin.Context) {
	var form request.CompList

	if err := c.ShouldBindQuery(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	compList, total, err := h.compS.GetComps(c, &form)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, &resp.RespList{List: compList, Total: total})
}

func (h *CompHandler) GetCompInfo(c *gin.Context) {

	compId := c.Query("id")
	if compId == "" {
		response.FailByErr(c, errors.New("id不能为空"))
		return
	}

	id, _ := strconv.ParseUint(compId, 10, 64)

	comp, err := h.compS.GetCompById(c, id)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, comp)

}

func (h *CompHandler) UpdateCompInfo(c *gin.Context) {
	var form request.UpdateCompInfo

	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

	comp, err := h.compS.UpdateComp(c, &form)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, comp)

}

func (h *CompHandler) AuditComp(c *gin.Context) {
	var form request.AuditComp
	if err := c.ShouldBindJSON(&form); err != nil {
		response.FailByErr(c, request.GetError(form, err))
		return
	}

}
