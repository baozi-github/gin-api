package router

import (
	api2 "gin-api/controller/api"
	"gin-api/middleware/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initApiRouter(r *gin.Engine) {
	// 使用中间件 SetLanguageMiddleware 设置请求语言
	apiRouter := r.Group("/api").Use(api.SetLanguageMiddleware(), api.CheckContentType(), api.CheckSign())
	{
		apiRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		login := api2.NewLogin()
		apiRouter.POST("/login", login.Login)
		apiRouter.POST("/register", login.Register)

		apiRouter.Use(api.CheckJwt())
		{
			// 用户模块
			user := api2.NewUser()
			apiRouter.POST("/user/info", user.Info)
		}
	}
}
