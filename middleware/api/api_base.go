package api

import (
	"gin-api/common"
	"gin-api/language"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CheckContentType 校验api请求的基础数据
func CheckContentType() gin.HandlerFunc {
	return func(context *gin.Context) {
		contentType := context.GetHeader("Content-Type")
		if contentType != "application/json" {
			log.Println("application error ", contentType)
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
		}

		// 请求头增加版本号
		version := context.GetHeader("Version")
		if version == "" {
			log.Println("version error ", contentType)
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
		}
		context.Next()
	}
}
