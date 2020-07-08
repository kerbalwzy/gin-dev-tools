package kerbalwzygo

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"unicode"
)

// Get the MD5 hash value of bytes
func BytesMD5Hash(data []byte) string {
	mac := hmac.New(md5.New, nil)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

// Get the MD5 hash value of string
func StringMD5Hash(data string) string {
	return BytesMD5Hash([]byte(data))
}

// Get the MD5 hash value of multi string
func MultiStringMD5Hash(data ...string) string {
	var temp string
	for _, item := range data {
		temp += item
	}
	return StringMD5Hash(temp)
}

// 检查字符串是否包含中文字符
func StringContainsHan(data string) bool {
	temp := []rune(data)
	for _, r := range temp {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

// 检查字符串是否包含空白字符
func StringContainsSpace(data string) bool {
	temp := []rune(data)
	for _, r := range temp {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

var ErrStartLargeThenEnd = errors.New("the 'start' index value is over then 'end'")
var ErrIndexOutOfRange = errors.New("the 'start' or 'end' index value is out of range")

// 安全的切割字符串: data原字符串, start切片起点, end切片终点+1
// Direct cutting of string may cause character scrambling, because some character could be 3 or 4 byte.
// Translate the 'string' into 'rune' before cutting could make safe
func SafeSliceString(data string, start, end int) (string, error) {
	if start > end {
		return "", ErrStartLargeThenEnd
	}
	if start < 0 || end > len(data) {
		return "", ErrIndexOutOfRange
	}
	return string([]rune(data)[start:end]), nil
}
