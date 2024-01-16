package osenv

import (
	"os"
)

// # core api 验证token core_api_token
// CORE_API_TOKEN=
func GetCoreApiToken() string {
	return os.Getenv("CORE_API_TOKEN")
}

// #core api url https://server.domain
// CORE_API_URL=
func GetCoreApiUrl() string {
	return os.Getenv("CORE_API_URL")
}
