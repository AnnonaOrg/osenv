package osenv

import (
	"os"
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
func GetMtprotoProxySecret() string {
	return os.Getenv("MTPROTO_PROXY_SECRET")
}
