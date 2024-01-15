package osenv

import (
	"os"
)

// # core api 验证token core_api_token
// CORE_API_TOKEN=
func GetCoreApiToken() string {
	return os.Getenv("CORE_API_TOKEN")
}
