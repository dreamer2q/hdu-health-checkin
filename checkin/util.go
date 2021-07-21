package checkin

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

const (
	base       = "https://api.hduhelp.com"
	baseUrl    = "https://api.hduhelp.com/base/"
	checkInUrl = baseUrl + "healthcheckin?sign=%s"
	infoUrl    = baseUrl + "person/info"
	dailyUrl   = baseUrl + "healthcheckin/info/daily"
	validUrl   = "https://api.hduhelp.com/token/validate"

	todayLeaveUrl   = base + "/workflow/dayoff/today"
	requestLeaveUrl = base + "/workflow/request?templateID=992ec670-5dfa-47ae-9ca3-1646c1eb4d0d"

	agent = "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/78.0.3904.62 XWEB/2759 MMWEBSDK/201201 Mobile Safari/537.36 MMWEBID/3149 MicroMessenger/7.0.22.1820(0x2700163B) Process/toolsmp WeChat/arm64 Weixin NetType/WIFI Language/zh_CN ABI/arm64"
)

func b64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func m5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func s1(data string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(data)))
}
