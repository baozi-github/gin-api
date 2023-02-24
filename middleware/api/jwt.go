package api

import (
	"gin-api/cache"
	"gin-api/common"
	"gin-api/language"
	"gin-api/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type JwtClaim struct {
	Uid          string
	ValidateDate string
}

func CheckJwt() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("token")
		if tokenString == "" {
			log.Println("token is empty")
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
			return
		}

		claims, err := util.ParseToken(tokenString)
		if err != nil {
			log.Println("token parse error", err.Error())
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
			return
		}

		uid := claims.Uid

		// 获取缓存中的token 与 传递token做对比 如果不一致提示token无效
		key := common.RedisUserToken
		cacheToken := cache.RedisGet(context, key)
		if cacheToken != tokenString {
			log.Println("token invalid", err.Error())
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
			return
		}

		log.Println("token parse success ", claims)
		context.Set("uid", uid)
		context.Next()
	}
}
