package resp

type RespList struct {
	Total int64       `form:"total,string" json:"total,string"`
	List  interface{} `form:"list" json:"list"`
}
