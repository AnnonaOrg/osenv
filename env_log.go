package osenv

import (
	"os"
)

// LogLevel
func GetLogLevel() string {
	return os.Getenv("LOG_LEVEL")
}
