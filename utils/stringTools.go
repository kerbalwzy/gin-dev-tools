package utils

import (
    "crypto/hmac"
    "crypto/md5"
    "encoding/hex"
    "errors"
    "unicode"
)

// Get the MD5 hash value of bytes
func BytesHash(data []byte) string {
    mac := hmac.New(md5.New, nil)
    mac.Write(data)
    return hex.EncodeToString(mac.Sum(nil))
}

// Get the MD5 hash value of string
func StringHash(data string) string {
    return BytesHash([]byte(data))
}

// Get the MD5 hash value of multi string
func MultiStringHash(data ...string) string {
    var temp string
    for _, item := range data {
        temp += item
    }
    return StringHash(temp)
}

// Check the string whether contains chinese char
func StringContainHan(data string) bool {
    temp := []rune(data)
    for _, r := range temp {
        if unicode.Is(unicode.Han, r) {
            return true
        }
    }
    return false
}

var ErrStartLargeThenEnd = errors.New("the 'start' index value is large then 'end'")
var ErrIndexOutOfRange = errors.New("the 'start' of 'end' index value is out of range")

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
