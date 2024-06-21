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
	retInt, err := strconv.ParseInt(retText, 10, 32)
	if err != nil {
		return 0
	}
	return retInt
}
func GetMtprotoProxySecret() string {
	return os.Getenv("MTPROTO_PROXY_SECRET")
}

// #Socks5
func IsEnableSocks5Proxy() bool {
	return strings.EqualFold("true", os.Getenv("SOCKS5_PROXY_ENABLE"))
}
func GetSocks5ProxyServer() string {
	return os.Getenv("SOCKS5_PROXY_SERVER")
}
func GetSocks5ProxyPort() string {
	return os.Getenv("SOCKS5_PROXY_PORT")
}
func GetSocks5ProxyPortInt() int64 {
	retText := os.Getenv("SOCKS5_PROXY_PORT")
	retInt, err := strconv.ParseInt(retText, 10, 32)
	if err != nil {
		return 0
	}
	return retInt
}
func GetSocks5ProxyUsername() string {
	return os.Getenv("SOCKS5_PROXY_USERNAME")
}
func GetSocks5ProxyPassword() string {
	return os.Getenv("SOCKS5_PROXY_PASSWORD")
}
