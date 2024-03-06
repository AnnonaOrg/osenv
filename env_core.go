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

// # core api 验证token (只读权限) core_api_token_readonly
// CORE_API_TOKEN_READONLY=
func GetCoreApiToken() string {
	return os.Getenv("CORE_API_TOKEN_READONLY")
}
