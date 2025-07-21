package osenv

import (
	"os"
	"strings"
)

// #支付方式
// #EPUSDT 标识
// PAYMENT_EPUSDT_API_NAME=USDT
func GetPaymentEPUSDTName() string {
	retText := os.Getenv("PAYMENT_EPUSDT_API_NAME")
	if len(retText) > 0 {
		return retText
	}
	return "USDT"
}
func IsEnablePaymentEPUSDT() bool {
	return strings.EqualFold("true", os.Getenv("PAYMENT_EPUSDT_API_ENABLE"))
}

// #EPUSDT 接口地址
// PAYMENT_EPUSDT_API_URL=http://127.0.0.1:8000
func GetPaymentEPUSDTUrl() string {
	return os.Getenv("PAYMENT_EPUSDT_API_URL")
}

// #EPUSDT 验证密钥
// PAYMENT_EPUSDT_API_KEY=
func GetPaymentEPUSDTKey() string {
	return os.Getenv("PAYMENT_EPUSDT_API_KEY")
}

// #
// #BTCPay 标识
// PAYMENT_BTCPAY_API_NAME=BTCPay
func GetPaymentBTCPayName() string {
	retText := os.Getenv("PAYMENT_BTCPAY_API_NAME")
	if len(retText) > 0 {
		return retText
	}
	return "BTCPay"
}
func IsEnablePaymentBTCPay() bool {
	return strings.EqualFold("true", os.Getenv("PAYMENT_BTCPAY_API_ENABLE"))
}

// #BTCPay 接口地址
// PAYMENT_BTCPAY_API_URL=https://mainnet.demo.btcpayserver.org
func GetPaymentBTCPayUrl() string {
	return os.Getenv("PAYMENT_BTCPAY_API_URL")
}

// #BTCPay 验证密钥
// PAYMENT_BTCPAY_API_KEY=
func GetPaymentBTCPayKey() string {
	return os.Getenv("PAYMENT_BTCPAY_API_KEY")
}

// # BTCPay StoreID
// PAYMENT_BTCPAY_API_STOREID =
func GetPaymentBTCPayStoreID() string {
	return os.Getenv("PAYMENT_BTCPAY_API_STOREID")
}

// # BTCPay WebhookSecret
// PAYMENT_BTCPAY_API_WEBHOOK_SECRET=
func GetPaymentBTCPayWebhookSecret() string {
	return os.Getenv("PAYMENT_BTCPAY_API_WEBHOOK_SECRET")
}

// #
// #CryptAPI 标识
// PAYMENT_CRYPTAPI_API_NAME=CryptAPI
func GetPaymentCryptAPIName() string {
	retText := os.Getenv("PAYMENT_CRYPTAPI_API_NAME")
	if len(retText) > 0 {
		return retText
	}
	return "CryptAPI"
}
func IsEnablePaymentCryptAPI() bool {
	return strings.EqualFold("true", os.Getenv("PAYMENT_CRYPTAPI_API_ENABLE"))
}

// #
// #易支付 标识
// PAYMENT_EASYPAY_API_NAME=易支付
func GetPaymentEasyPayName() string {
	retText := os.Getenv("PAYMENT_EASYPAY_API_NAME")
	if len(retText) > 0 {
		return retText
	}
	return "EasyPay"
}
func IsEnablePaymentEasyPay() bool {
	return strings.EqualFold("true", os.Getenv("PAYMENT_EASYPAY_API_ENABLE"))
}

// #易支付 接口地址
// PAYMENT_EASYPAY_API_URL=
func GetPaymentEasyPayUrl() string {
	return os.Getenv("PAYMENT_EASYPAY_API_URL")
}

// #易支付 商户id
// PAYMENT_EASYPAY_API_PID=
func GetPaymentEasyPayPid() string {
	return os.Getenv("PAYMENT_EASYPAY_API_PID")
}

// #易支付 验证密钥
// PAYMENT_EASYPAY_API_KEY=
func GetPaymentEasyPayKey() string {
	return os.Getenv("PAYMENT_EASYPAY_API_KEY")
}

// #易支付 接口请求编码格式 JsonEncode URLValuesEncode
// PAYMENT_EASYPAY_REQUEST_BODY_ENCODE
func GetPaymentEasyRequstBodyEncode() string {
	return os.Getenv("PAYMENT_EASYPAY_REQUEST_BODY_ENCODE")
}

// #易支付 设备类型
// PAYMENT_EASYPAY_API_DEVICE=
func GetPaymentEasyPayDevice() string {
	return os.Getenv("PAYMENT_EASYPAY_API_DEVICE")
}

// #Boxcoin 标识
func GetPaymentBoxcoinName() string {
	retText := os.Getenv("PAYMENT_BOXCOIN_API_NAME")
	if len(retText) > 0 {
		return retText
	}
	return "Boxcoin"
}
func IsEnablePaymentBoxcoin() bool {
	return strings.EqualFold("true", os.Getenv("PAYMENT_BOXCOIN_API_ENABLE"))
}

// #Boxcoin 接口地址 YOUR-DOMAIN/boxcoin/api.php
func GetPaymentBoxcoinUrl() string {
	return os.Getenv("PAYMENT_BOXCOIN_API_URL")
}

// #Boxcoin api 验证密钥
func GetPaymentBoxcoinKey() string {
	return os.Getenv("PAYMENT_BOXCOIN_API_KEY")
}

// #Boxcoin webhook 验证密钥
func GetPaymentBoxcoinWebhookKey() string {
	return os.Getenv("PAYMENT_BOXCOIN_WEBHOOK_KEY")
}

func GetPaymentUmiPayName() string {
	retText := os.Getenv("PAYMENT_UMIPAY_API_NAME")
	if len(retText) > 0 {
		return retText
	}
	return "UmiPay"
}
func IsEnablePaymentUmiPay() bool {
	return strings.EqualFold("true", os.Getenv("PAYMENT_UMIPAY_API_ENABLE"))
}
func GetPaymentUmiPayUrl() string {
	return os.Getenv("PAYMENT_UMIPAY_API_URL")
}

// 商户id
func GetPaymentUmiPayPid() string {
	return os.Getenv("PAYMENT_UMIPAY_API_PID")
}

// 验证密钥
func GetPaymentUmiPayKey() string {
	return os.Getenv("PAYMENT_UMIPAY_API_KEY")
}
