package request

// 注册结构

type NewComp struct {
	Title    string `form:"title" json:"title" binding:"required"`
	Info     string `form:"info" json:"info"`
	CreateId uint64 `form:"create_id" json:"create_id,string"`
	Deploy   string `form:"deploy" json:"deploy" binding:"required"`
	//Types      int    `form:"types" json:"types,string"`
	PreviewUrl string `form:"preview_url" json:"preview_url" binding:"required"`
	Url        string `form:"url" json:"url" binding:"required"`
	Row        int    `form:"row" json:"row" binding:"required"`
	Column     int    `form:"column" json:"column" binding:"required"`
}

func (newComp NewComp) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Title.required":      "组件名称不能为空",
		"Deploy.required":     "组件部署方式不能为空",
		"PreviewUrl.required": "组件预览地址不能为空",
		"Url.required":        "组件访问地址不能为空",
		"Column.required":     "默认列数不能为空",
		"Row.required":        "默认行数不能为空",
	}
}

// CompList 获取组件列表结构
type CompList struct {
	Name     string `form:"name" json:"name"`
	CreateId uint   `form:"createId,string" json:"createId,string"`
	ID       uint   `form:"id,string" json:"id,string"`
	PageDto
}

// UpdateCompInfo 更新组件结构
type UpdateCompInfo struct {
	ID         uint64 `form:"id" json:"id,string"` // 主键
	Title      string `form:"title" json:"title" binding:"required"`
	Info       string `form:"info" json:"info"`
	CreateId   uint64 `form:"create_id" json:"create_id,string"`
	Deploy     string `form:"deploy" json:"deploy" binding:"required"`
	PreviewUrl string `form:"preview_url" json:"preview_url" binding:"required"`
	Url        string `form:"url" json:"url" binding:"required"`
	Row        int    `form:"row" json:"row" binding:"required"`
	Column     int    `form:"column" json:"column" binding:"required"`
}

func (updateCompInfo UpdateCompInfo) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Title.required":      "组件名称不能为空",
		"Deploy.required":     "组件部署方式不能为空",
		"PreviewUrl.required": "组件预览地址不能为空",
		"Url.required":        "组件访问地址不能为空",
		"Column.required":     "默认列数不能为空",
		"Row.required":        "默认行数不能为空",
	}
}

type AuditComp struct {
	ID       uint64 `form:"id" json:"id,string"`               // 组件
	IsList   int    `form:"is_list" json:"is_list"`            // 审核状态
	Msg      string `form:"msg" json:"msg"`                    // 审核意见
	CreateId uint64 `form:"create_id" json:"create_id,string"` // 审核人id
}

func (auditComp AuditComp) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"ID.required":       "组件id不能为空",
		"IsList.required":   "审核状态不能为空",
		"Msg.required":      "审核意见不能为空",
		"CreateId.required": "审核人id不能为空",
	}
}

type NewPage struct {
	CreateId uint64      `form:"create_id" json:"create_id,string"` // 创建人
	Data     interface{} `form:"data" json:"data"`                  // 页面数据
	Info     string      `form:"info" json:"info"`                  // 页面信息
	Title    string      `form:"info" json:"title"`                 // 页面标题
}

func (newPAge NewPage) GetMessage() ValidatorMessages {
	return ValidatorMessages{
		"CreateId.required": "组件id不能为空",
		"Data.required":     "审核状态不能为空",
	}

}
