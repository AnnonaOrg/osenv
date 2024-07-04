package osenv

import (
	"os"
)

// Bot token
func GetAppTelegramBotToken() string {
	return os.Getenv("APP_TELEGRAM_BOT_TOKEN")
}

// app api id
func GetAppTelegramApiID() string {
	return os.Getenv("APP_TELEGRAM_API_ID")
}

// app api hash
func GetAppTelegramApiHash() string {
	return os.Getenv("APP_TELEGRAM_API_HASH")
}

// app user phone number
func GetAppTelegramUserPhoneNumber() string {
	return os.Getenv("APP_TELEGRAM_USER_PHONE_NUMBER")
}

// app user id
func GetAppTelegramUserID() string {
	return os.Getenv("APP_TELEGRAM_USER_ID")
}

func IsTDlibSimpleMessage() bool {
	return strings.EqualFold("true", os.Getenv("TDLIB_SIMPLE_MESSAGE_ENABLE"))
}
