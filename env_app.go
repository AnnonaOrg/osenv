package osenv

import (
	"os"
	"strings"
)

// app api id
func GetAppTelegramApiID() string {
	return os.Getenv("APP_TELEGRAM_API_ID")
}

// app api hash
func GetAppTelegramApiHash() string {
	return os.Getenv("APP_TELEGRAM_API_HASH")
}

func IsTDlibSimpleMessage() bool {
	return strings.EqualFold("true", os.Getenv("TDLIB_SIMPLE_MESSAGE_ENABLE"))
}
