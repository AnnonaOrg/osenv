package osenv

import (
	"os"
)

// #fenxiao api url https://server.domain
// FENXIAO_API_URL=
func GetFenxiaoApiUrl() string {
	return os.Getenv("FENXIAO_API_URL")
}
