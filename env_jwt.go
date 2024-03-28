package osenv

import (
	"math/rand"
	"os"
	"sync"
)

// #JWT Secret
// JWT_SECRET=
var JWTSecret string
var once sync.Once

func GetJWTSecret() string {
	item := os.Getenv("JWT_SECRET")
	if len(item) > 0 {
		return item
	}
	return GenerateJWTSecret()
}

func GenerateJWTSecret() string {
	if len(JWTSecret) > 0 {
		return JWTSecret
	} else {
		once.Do(func() {
			JWTSecret = GenerateRandomString(10)
		})
		return JWTSecret
	}
}

func GenerateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
