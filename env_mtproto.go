package osenv

import (
	"os"
	"strconv"
	"strings"
)

// #MtprotoProxy
func IsEnableMtprotoProxy() bool {
	return strings.EqualFold("true", os.Getenv("MTPROTO_PROXY_ENABLE"))
}
func GetMtprotoProxyServer() string {
	return os.Getenv("MTPROTO_PROXY_SERVER")
}
func GetMtprotoProxyPort() string {
	return os.Getenv("MTPROTO_PROXY_PORT")
}
func GetMtprotoProxyPortInt() int64 {
	// return os.Getenv("MTPROTO_PROXY_PORT")
	retText := os.Getenv("MTPROTO_PROXY_PORT")
	retInt, err := strconv.ParseInt(retText, 10, 64)
	if err != nil {
		return 0
	}
	return retInt
}
func GetMtprotoProxySecret() string {
	return os.Getenv("MTPROTO_PROXY_SECRET")
}
