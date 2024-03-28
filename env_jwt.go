package osenv

import (
	"os"
)

// #JWT Secret
// JWT_SECRET=
func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
