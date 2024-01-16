package osenv

import (
	"os"
)

// #平台类型 tele
// #PLATFORM_TYPE_TELE
// PLATFORM_TYPE_TELE=tele
func GetPlatformType() string {
	return os.Getenv("PLATFORM_TYPE_TELE")
}
