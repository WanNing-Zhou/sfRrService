package app

import (
	"github.com/gin-gonic/gin"
	cErr "github.com/jassue/gin-wire/app/pkg/error"
	"github.com/jassue/gin-wire/app/pkg/request"
	"github.com/jassue/gin-wire/app/pkg/resp"
	"github.com/jassue/gin-wire/app/pkg/response"
	"github.com/jassue/gin-wire/app/service"
	"go.uber.org/zap"
	"strconv"
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

	err := c.ShouldBindJSON(&form)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	//// 获取到body数据
	//body, _ := io.ReadAll(c.Request.Body)
	//fmt.Println(body)
	//// 将body转换为map
	//var data map[string]map[string]interface{}
	//_ = json.Unmarshal(body, &data)
	//// 获取到data数据
	//pageData := data["data"]
	//// 将pageData转换为json
	//marshal, _ := json.Marshal(pageData)

	//form.Data = string(marshal)

	//fmt.Println("marshal", string(marshal))

	form.CreateId, err = strconv.ParseUint(c.Keys["id"].(string), 10, 64)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	err = h.pageS.Create(c, &form)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	successMap := map[string]string{
		"msg": "success",
	}

	response.Success(c, successMap)
}

func (h *PageHandler) GetPageList(c *gin.Context) {
	id, err := strconv.ParseUint(c.Keys["id"].(string), 10, 64)
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	param := &request.GetPages{
		CreateId: id,
		PageDto: request.PageDto{
			PageSize: 100,
			Page:     1,
		},
	}

	pages, i, err := h.pageS.GetPages(c, param)
	if err != nil {
		response.FailByErr(c, err)
		return
	}
	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, &resp.RespList{List: pages, Total: i})

}

func (h *PageHandler) Info(c *gin.Context) {
	id := c.Query("id")
	//fmt.Println("id", )
	if id == "" {
		response.FailByErr(c, cErr.ValidateErr("id不存在"))
		return
	}

	page, err := h.pageS.FindById(c, id)

	if err != nil {
		response.FailByErr(c, err)
		return
	}

	response.Success(c, page)
}
