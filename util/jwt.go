package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

type Claims struct {
	Uid string `json:"uid"` // 用户id
	jwt.StandardClaims
}

// GenerateToken 生成token
// uid string 用户id
// validDate string 有效期 小时
func GenerateToken(uid string, validDate string) string {
	// 默认有效期 30天
	tokenExpireDuration := time.Hour * 720
	if validDate != "" {
		validDateHour, err := strconv.Atoi(validDate)
		if err == nil {
			tokenExpireDuration = time.Duration(validDateHour) * time.Hour
		}
	}

	expirationTime := time.Now().Add(tokenExpireDuration)
	log.Println("token 过期时间", expirationTime)
	jwtSecret := viper.GetString("jwt.secret")
	claim := Claims{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwtSecretByte := []byte(jwtSecret)
	tokenString, err := token.SignedString(jwtSecretByte)
	if err != nil {
		log.Println("token 生成失败 ", err.Error())
		return ""
	}

	return tokenString
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	jwtSecret := viper.GetString("jwt.secret")
	jwtSecretByte := []byte(jwtSecret)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing methon")
		}
		return jwtSecretByte, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			// 判断token是否过期
			return claims, nil
		}
	}

	return nil, err
}
