package common

import (
	"encoding/json"
	"gin-api/language"
	"gin-api/util/rsa"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

type ResponseData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ReturnSuccessData 组装请求成功的数据返回值
func ReturnSuccessData(data interface{}, ctx *gin.Context) map[string]string {
	code := language.GetCodeByLang(language.SUCCESS)
	msg := language.GetMessageByLang(language.SUCCESS, ctx)
	res := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	dataStr := Encrypt(res)
	return map[string]string{
		"data": dataStr,
	}
}

// ReturnFailData 组装请求失败的返回值
func ReturnFailData(dictionaryStr string, message string, data interface{}, ctx *gin.Context) map[string]string {
	code := language.GetCodeByLang(dictionaryStr)
	msg := language.GetMessageByLang(dictionaryStr, ctx)
	if code == language.SystemCode[language.ParamError] && message != "" {
		msg = message
	}

	res := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	dataStr := Encrypt(res)
	return map[string]string{
		"data": dataStr,
	}
}

// Encrypt 私钥加密
func Encrypt(data ResponseData) string {
	strByte, err := json.Marshal(data)
	log.Println("encrypt data ", strByte)
	if err != nil {
		log.Println("数据加密失败01", err.Error())
		return ""
	}

	publicKey := viper.GetString("rsa_request.V3.public_key")
	// 公钥加密
	encryptStr := rsa.Encrypt(strByte, publicKey)
	if encryptStr == "" {
		log.Println("数据加密失败")
	}

	// 测试解密
	privateKey := viper.GetString("rsa_request.V3.private_key")
	decryptStr := rsa.RsaDecrypt(encryptStr, privateKey)
	log.Println("decryptStr", decryptStr)

	return encryptStr
}
