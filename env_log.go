package osenv

import (
	"os"
)

// LogLevel panic,fatal,error,warn,warning,info,debug,trace
func GetLogLevel() string {
	if retText := os.Getenv("LOG_LEVEL"); len(retText) > 0 {
		return retText
	}
	return "info"
}
