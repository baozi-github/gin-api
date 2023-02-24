package api

// LoginData 传入非字符串的时候会报错
type LoginData struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var LoginDataErrMsg = map[string]string{
	"Mobile.required":   "手机号为必填项",
	"Password.required": "密码为必填项",
}
