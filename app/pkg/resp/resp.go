package resp

type RespList struct {
	Total int64       `form:"total" json:"total"`
	List  interface{} `form:"list" json:"list"`
}

type SuccessData struct {
	Message string `form:"message" json:"message"`
}
