package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

// Application 定义全局变量
type Application struct {
	DB  *gorm.DB
	RDB *redis.Client
}

var App Application
