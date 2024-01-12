package osenv

import (
	"os"
)

// Bot token
func GetAppTelegramBotToken() string {
	return os.Getenv("APP_TELEGRAM_BOT_TOKEN")
}

// apiIdRaw = os.Getenv("API_ID")
// apiHash  = os.Getenv("API_HASH")
func GetAppTelegramApiID() string {
	return os.Getenv("APP_TELEGRAM_API_ID")
}
func GetAppTelegramApiHash() string {
	return os.Getenv("APP_TELEGRAM_API_HASH")
}
