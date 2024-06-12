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

// bot apiproxy url
func GetBotTelegramAPIProxyURL() string {
	return os.Getenv("BOT_TELEGRAM_API_PROXY_URL")
}

// 告警通知群id: 接收反馈消息的Chat id
func GetBotReportChatID() string {
	return os.Getenv("BOT_REPORT_CHAT_ID")
}

// welecome msg
func GetBotWelcomeMsg() string {
	return os.Getenv("BOT_WELCOME_MSG")
}

// 文本消息尾巴
func GetBotCaptionExt() string {
	return os.Getenv("BOT_CAPTION_EXT")
}
