package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	initApiRouter(r)
}
