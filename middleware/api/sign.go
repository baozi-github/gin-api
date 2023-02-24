package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gin-api/common"
	"gin-api/language"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type MustData struct {
	Sign string `json:"sign" form:"sign" binding:"required"`
	Time string `json:"time" form:"time" binding:"required"`
}

func CheckSign() gin.HandlerFunc {
	return func(context *gin.Context) {
		isOpen := viper.GetString("sign.open")
		// 如果签名是未开启状态
		if isOpen == "0" {
			context.Next()
			return
		}

		var data map[string]interface{}
		var mustData MustData
		// 固定参数校验
		if err := context.ShouldBindBodyWith(&mustData, binding.JSON); err != nil {
			log.Println("sign must data error ", err.Error())
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
			return
		}

		timeSecondNow := time.Now().Unix()
		timeParam, err := strconv.Atoi(mustData.Time)
		if err != nil {
			log.Println("sign time parse error", err.Error(), mustData)
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
			return
		}

		// 请求时间超出15秒 错误  请求提前5秒 返回
		if timeSecondNow-int64(timeParam) > 15 || int64(timeParam)-timeSecondNow > 5 {
			log.Println("sign time out error ", timeSecondNow, timeParam)
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
			return
		}

		// 绑定参数 失败则直接返回签名错误
		if err := context.ShouldBindBodyWith(&data, binding.JSON); err != nil {
			log.Println("sign bind data error ", err.Error())
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
			return
		}

		// 获取传递的sign 不存在则直接返回签名错误
		paramSign, ok := data["sign"]
		if !ok {
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
		}

		// sign字段必须是字符串 如果是其他类型则转成字符串
		paramSignSting := fmt.Sprintf("%v", paramSign)

		// 根据传递参数 本地生成sign
		sign := createSign(data)

		// 自己生成的sign与用户传递sign做对比
		if sign == "" || sign != paramSignSting {
			res := common.ReturnFailData(language.TokenValidate, "", map[string]string{}, context)
			context.JSON(http.StatusOK, res)
			context.Abort()
		}
		log.Println("data == ", data)
		log.Println("sign == ", sign)

		context.Next()
	}
}

// 生成签名
func createSign(data map[string]interface{}) string {
	// 从配置文件读取
	salt := viper.GetString("sign.salt")
	str := ""
	// 需要签名的字段
	var keyArr []string
	for key := range data {
		// 跳过签名
		if key == "sign" {
			continue
		}
		keyArr = append(keyArr, key)
	}

	// 排序
	sort.Strings(keyArr)

	// 拼接字符串 并将非字符串类型转为字符串
	for i := 0; i < len(keyArr); i++ {
		// 初始值
		startValue := data[keyArr[i]]
		var stringValue string
		var formatString string
		if i == 0 {
			formatString = "%v=%v"
		} else {
			formatString = "&%v=%v"
		}

		switch startValue.(type) {
		case string:
			stringValue = fmt.Sprintf(formatString, keyArr[i], startValue)
		default:
			startValueByte, err := json.Marshal(startValue)
			if err != nil {
				log.Println("sign json encode error", startValue)
				return ""
			}
			stringValue = string(startValueByte)
			stringValue = fmt.Sprintf(formatString, keyArr[i], stringValue)
		}
		str = str + stringValue
	}

	str = str + salt
	// md5加密
	s := md5.New()
	s.Write([]byte(str))
	beginSign := hex.EncodeToString(s.Sum(nil))

	// 再次加密
	index := []int{1, 2, 4, 6, 7, 9, 11, 12, 14, 15}
	endStr := ""
	for _, v := range index {
		endStr = endStr + string(beginSign[v])
	}

	ss := md5.New()
	ss.Write([]byte(str))
	endSign := hex.EncodeToString(ss.Sum(nil))

	return endSign
}
