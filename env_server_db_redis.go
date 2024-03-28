package osenv

import (
	"os"
	"strconv"
)

// server db redis address
// SERVER_DB_REDIS_ADDRESS 127.0.0.1:6379
func GetServerDbRedisAddress() string {
	return os.Getenv("SERVER_DB_REDIS_ADDRESS")
}

// server db redis password
// SERVER_DB_REDIS_PASSWORD Redis密码
func GetServerDbRedisPassword() string {
	return os.Getenv("SERVER_DB_REDIS_PASSWORD")
}

// server db redis Redis库
// SERVER_DB_REDIS_DB 0 1 2 3
func GetServerDbRedisDB() int {
	v := os.Getenv("SERVER_DB_REDIS_DB")
	if v == "" {
		return 0
	}
	if vi, err := strconv.Atoi(v); err != nil {
		return 0
	} else {
		vi
	}
}
