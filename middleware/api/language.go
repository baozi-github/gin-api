package api

import (
	"gin-api/language"
	"github.com/gin-gonic/gin"
)

// SetLanguageMiddleware 中间件 获取用户选择的语言
func SetLanguageMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		langList := language.GetAllowLangArr()
		langHeader := context.GetHeader("lang")
		lang := language.GetDefaultLang()

		for _, item := range langList {
			if item == langHeader {
				lang = langHeader
			}
		}
		// 设置默认语言
		setRequestLang(lang, context)
		context.Next()
	}

}

// 设置语言
func setRequestLang(lang string, ctx *gin.Context) {
	ctx.Set("lang", lang)
}
