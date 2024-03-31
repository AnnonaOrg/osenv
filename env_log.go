package osenv

import (
	"os"
)

// LogLevel panic,fatal,error,warn,warning,info,debug,trace
func GetLogLevel() string {
	retText := os.Getenv("LOG_LEVEL")
	if len(retText) > 0 {
		return retText
	}
	return "info"
}
