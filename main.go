package main

import (
	"bytes"
	"context"
	"gin-api/config"
	"gin-api/crontab"
	"gin-api/global"
	"gin-api/router"
	localValidate "gin-api/validate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

// 加载配置文件
func init() {
	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewBuffer(config.Config)); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("no config file")
		} else {
			log.Println("read config error", err.Error())
		}
		log.Fatal(err)
	}
}

// 初始化redis
func initMysql() {
	var err error
	dbUser := viper.GetString("mysql.dbUser")
	dbPassword := viper.GetString("mysql.dbPassword")
	dbName := viper.GetString("mysql.dbName")
	dbHost := viper.GetString("mysql.dbHost")
	dbPort := viper.GetString("mysql.dbPort")

	dsn := dbUser + ":" + dbPassword + "@tcp" + "(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	log.Println("dsn = ", dsn)

	global.App.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("fail to connect to mysql", err)
		return
	}
}

// 初始化redis
func initRedis() {
	addr := viper.GetString("redis.addr")
	password := viper.GetString("redis.password")
	db := viper.GetString("redis.db")
	dbInt, _ := strconv.Atoi(db)
	global.App.RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbInt,
	})
	result := global.App.RDB.Ping(context.Background())
	log.Println("redis ping:", result.Val())

	if result.Val() != "PONG" {
		log.Println("redis connect error")
	}
}

func main() {
	r := gin.Default()
	// 初始化数据库
	initMysql()
	// 初始化redis
	initRedis()
	// 注册验证规则
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = validate.RegisterValidation("is_string", localValidate.IsString)
	} else {
		log.Println("validate bind err ")
	}
	// 注册定时器
	crontab.InitCrontab()

	// 关闭mysql
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			log.Println("mysql close err ", err.Error())
		}
	}(global.App.DB)

	// 获取配置文件
	port := viper.GetString("app.port")
	log.Println("监听端口", "http://0.0.0.0:"+port)
	// 初始化路由
	router.InitRouter(r)
	// 启动服务
	err := r.Run(":" + port)
	if err != nil {
		return
	}
}
