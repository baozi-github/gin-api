package validate

import (
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"strings"
)

// GetErrorMsg 获取自定义错误信息
func GetErrorMsg(errorMessageMap map[string]string, err error) string {
	// 没有传递任何参数时
	if err == io.EOF {
		return "参数绑定错误"
	}

	// 传递的参数类型与规定参数类型不一致时
	errString := err.Error()
	if strings.Contains(errString, "unmarshal") {
		log.Println("参数类型错误")
		return "参数类型错误"
	}

	// 获取错误信息
	for _, v := range err.(validator.ValidationErrors) {
		key := v.Field() + "." + v.Tag()
		if message, exist := errorMessageMap[key]; exist {
			return message
		}
		return v.Error()
	}
	return "参数错误"
}

// IsString 判断参数是否是字符串 注册方法 is_string
func IsString(v validator.FieldLevel) bool {
	field := v.Field().Interface()
	switch field.(type) {
	case string:
		return true
	default:
		return false
	}
}
