package api

import (
	"gin-api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
}

func NewUser() User {
	return User{}
}

func (u *User) Info(ctx *gin.Context) {
	// 参数校验
	res := common.ReturnSuccessData(map[string]string{}, ctx)
	ctx.JSON(http.StatusOK, res)
	return
}
