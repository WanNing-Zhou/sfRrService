package request

// 注册结构

type NewComp struct {
	Title      string `form:"title" json:"title" binding:"required"`
	Info       string `form:"info" json:"info"`
	CreateId   uint64 `form:"create_id" json:"create_id,string"`
	Deploy     string `form:"deploy" json:"deploy" binding:"required"`
	Types      int    `form:"types" json:"types,string"`
	PreviewUrl string `form:"previewUrl" json:"previewUrl" binding:"required"`
	Url        string `form:"url" json:"url" binding:"required"`
}

func (newComp NewComp) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Title.required":      "组件名称不能为空",
		"Deploy.required":     "组件部署方式不能为空",
		"PreviewUrl.required": "组件预览地址不能为空",
		"Url.required":        "组件访问地址不能为空",
	}
}
