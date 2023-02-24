package api

import (
	"gin-api/cache"
	"gin-api/common"
	"gin-api/language"
	"gin-api/models"
	"gin-api/util"
	"gin-api/validate"
	validateapi "gin-api/validate/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

type Login struct {
}

type registerJson struct {
	Mobile   string `json:"mobile" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

// NewLogin 为什么要重新定义一个函数出来
func NewLogin() Login {
	return Login{}
}

// Login 登录
func (h *Login) Login(ctx *gin.Context) {
	var loginData validateapi.LoginData
	// 参数错误
	if err := ctx.ShouldBindBodyWith(&loginData, binding.JSON); err != nil {
		message := validate.GetErrorMsg(validateapi.LoginDataErrMsg, err)
		res := common.ReturnFailData(language.ParamError, message, map[string]string{}, ctx)
		ctx.JSON(http.StatusOK, res)
		return
	}

	// 打印数据
	res := common.ReturnSuccessData(map[string]string{}, ctx)
	ctx.JSON(http.StatusOK, res)
	return
}

// Register 注册
func (h *Login) Register(ctx *gin.Context) {
	var registerData registerJson
	// 参数错误
	if err := ctx.ShouldBind(&registerData); err != nil {
		res := common.ReturnFailData(language.ParamError, "", map[string]string{}, ctx)
		ctx.JSON(http.StatusOK, res)
		return
	}

	id := 1
	user, err := models.GetUser(id)
	if err != nil {
		res := common.ReturnFailData(language.ParamError, "", map[string]string{}, ctx)
		ctx.JSON(http.StatusOK, res)
		return
	}

	token := util.GenerateToken("1", "")
	if token == "" {
		res := common.ReturnFailData(language.ParamError, "", map[string]string{}, ctx)
		ctx.JSON(http.StatusOK, res)
		return
	}

	key := common.RedisUserToken + "1"
	flag := cache.RedisSet(ctx, key, token, time.Duration(720)*time.Hour)
	if !flag {
		res := common.ReturnFailData(language.ParamError, "", map[string]string{}, ctx)
		ctx.JSON(http.StatusOK, res)
		return
	}

	// 打印数据
	res := common.ReturnSuccessData(user, ctx)
	ctx.JSON(http.StatusOK, res)
	return
}
