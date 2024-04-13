package request

// 注册结构

type Register struct {
	Name string `form:"name" json:"name" binding:"required"`
	//Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Auth     int    `form:"auth" json:"auth"`
}

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required": "用户名称不能为空",
		//"Mobile.required":   "手机号码不能为空",
		//"Mobile.mobile":     "手机号码格式不正确",
		"Password.required": "用户密码不能为空",
		"Email.required":    "邮箱不能为空",
	}
}

// Login 登陆结构
type Login struct {
	//Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Mobile.required":   "手机号码不能为空",
		"Mobile.mobile":     "手机号码格式不正确",
		"Password.required": "用户密码不能为空",
	}
}

// Info 用户信息结构
type Info struct {
	ID   uint64 `form:"id" json:"id,string" binding:"required"` // id
	Name string `form:"name" json:"name" binding:"required"`
	//Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	//Password string `form:"password" json:"password" binding:"required"`
	Email        string `form:"email" json:"email" binding:"email"` // 邮箱
	Avatar       string `form:"avatar" json:"avatar"`               // 头像
	Mobile       string `form:"phone" json:"Mobile"`                // 电话
	Introduction string `form:"introduction" json:"introduction"`   // 简介
}

func (setInfo Info) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"ID.required":   "用户ID不能为空",
		"Name.required": "用户名称不能为空",
		//"Mobile.required":   "手机号码不能为空",
		//"Mobile.mobile":     "手机号码格式不正确",
		//"Password.required": "用户密码不能为空",
		"Email.required": "邮箱不能为空",
	}
}

// Password 用户密码
type Password struct {
	ID          uint64 `form:"id" json:"id,string" binding:"required"`
	OldPassword string `form:"oldPassword" json:"oldPassword" binding:"required"` // id
	NewPassword string `form:"newPassword" json:"newPassword" binding:"required"`
}

func (setInfo Password) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"ID.required":          "用户ID不能为空",
		"OldPassword.required": "旧密码不能为空",
		"NewPassword.required": "新密码不能为空",
	}
}

// GetUsers 获取用户结构
type GetUsers struct {
	Name  string `form:"name" json:"name"`
	ID    uint64 `form:"id" json:"id,string"`
	Email string `form:"email" json:"email"`
	PageDto
}

type GetUserInfo struct {
	ID string `form:"id" json:"id,string"`
}

type DeleteUser struct {
	ID uint64 `form:"id" json:"id,string"`
}
