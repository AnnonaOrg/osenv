package osenv

import (
	"os"
)

func GetCopyrightLicenseCode() string {
	return os.Getenv("COPYRIGHT_LICENSE_CODE")
}

// // true 授权验证通过
// func CheckCopyrightLicenseCode() bool {
// 	item := GetCopyrightLicenseCode()
// 	if len(item) > 0 {
// 		return CmpCopyrightLicenseCode(item)
// 	} else {
// 		return false
// 	}
// }
// func CmpCopyrightLicenseCode(copyrithtLicenseCode string) bool {
// 	if copyrithtLicenseCode == "https://t.me/umfaka" {
// 		return true
// 	}
// 	return false
// }
