package osenv

import (
	"fmt"
	"os"
	"strings"
)

// manager id
func GetBotManagerID() string {
	return os.Getenv("BOT_MANAGER_ID")
}
func IsBotManagerID(id int64) bool {
	return strings.EqualFold(
		GetBotManagerID(),
		fmt.Sprintf("%d", id),
	)
}
func IsBotManagerIDStr(idStr string) bool {
	return strings.EqualFold(
		GetBotManagerID(),
		idStr,
	)
}

// Bot token
func GetBotTelegramToken() string {
	return os.Getenv("BOT_TELEGRAM_TOKEN")
}

// bot webhook url
func GetBotTelegramWebhookURL() string {
	return os.Getenv("BOT_TELEGRAM_WEBHOOK_URL")
}
func GetBotAdminTelegramWebhookURL() string {
	return os.Getenv("BOT_ADMIN_TELEGRAM_WEBHOOK_URL")
}

// bot apiproxy url
func GetBotTelegramAPIProxyURL() string {
	return os.Getenv("BOT_TELEGRAM_API_PROXY_URL")
}
