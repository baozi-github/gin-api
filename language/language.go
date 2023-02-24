package language

import "github.com/gin-gonic/gin"

// 设置允许的语言类型
var allowAllArr = []string{"cn", "en"}

// 设置默认的语言类型
var defaultLang = "cn"

// 定义 语言类型对应的map
var langMap = map[string]map[string]string{
	"cn": langCn,
	"en": langEn,
}

// 定义返回值常量
const (
	SUCCESS          = "SUCCESS"
	ParamError       = "FailParamError"
	TokenValidate    = "TokenValidate"
	ContentTypeError = "ContentTypeError"
)

// SystemCode 定义返回code
var SystemCode = map[string]string{
	SUCCESS:          "0",
	ParamError:       "100010",
	TokenValidate:    "100011",
	ContentTypeError: "100012",
}

// 定义中文
var langCn = map[string]string{
	SUCCESS:          "成功",
	ParamError:       "参数错误",
	TokenValidate:    "登录凭证无效",
	ContentTypeError: "Content-Type无效",
}

// 定义英文
var langEn = map[string]string{
	SUCCESS:          "success",
	ParamError:       "param error",
	TokenValidate:    "token error",
	ContentTypeError: "Content-Type error",
}

// GetMessageByLang 根据语言类型获取message
func GetMessageByLang(dictionary string, ctx *gin.Context) string {
	lang := ctx.GetString("lang")
	langStr, ok := langMap[lang]
	if ok {
		return langStr[dictionary]
	}

	return "语言文件不完整"
}

// GetCodeByLang 获取消息码
func GetCodeByLang(dictionary string) string {
	code := SystemCode[dictionary]
	return code
}

// GetDefaultLang 获取默认的语言类型
func GetDefaultLang() string {
	return defaultLang
}

// GetAllowLangArr 获取允许的语言类型
func GetAllowLangArr() []string {
	return allowAllArr
}
