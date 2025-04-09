package osenv

import (
	"os"
)

// TRONGRID_API_KEY
func GetTrongridApikey() string {
	return os.Getenv("TRONGRID_API_KEY")
}
