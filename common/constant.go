package common

const (
	// RedisKeyVersion redis key 版本号
	RedisKeyVersion          = "v1_"
	REDIS_DEFAULT_EXPIRAYION = 0 // 默认缓存时间 0 用不过期
	RedisUserToken           = RedisKeyVersion + "user_token_"
)
