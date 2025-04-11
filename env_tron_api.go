package osenv

import (
	"os"
)

// TRONGRID_API_KEY  通过 https://www.trongrid.io 注册获取
func GetTrongridApikey() string {
	return os.Getenv("TRONGRID_API_KEY")
}

// TRON-PRO-API-KEY https://tronscan.org/#/myaccount/apiKeys/
func GetTronScanOrgApikey() string {
	return os.Getenv("TRONSCAN_ORG_API_KEY")
}
