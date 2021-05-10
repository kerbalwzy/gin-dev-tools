package kerbalwzygo

import (
	"time"
)

// Set time zone at UTC/GMT+08:00,  China Standard Time UTC+8:00
var BJS = time.FixedZone("BJS", 8*3600)

func NowTimestamp() int64 {
	return time.Now().Unix()
}

func Timestamp2Datetime(timestamp int64, tz *time.Location) time.Time {
	return time.Unix(timestamp, 0).In(tz)
}

func Timestamp2DatetimeStr(timestamp int64, tz *time.Location) string {
	return Timestamp2Datetime(timestamp, tz).Format("2006-01-02 15:04:05")
}

func BJSNowDatetimeStr() string {
	return time.Now().In(BJS).Format("2006-01-02 15:04:05")
}

func UTCNowDatetimeStr() string {
	return time.Now().In(time.UTC).Format("2006-01-02 15:04:05")
}

func BJSTodayDateStr() string {
	return time.Now().In(BJS).Format("2006-01-02")
}

func UTCTodayDateStr() string {
	return time.Now().In(time.UTC).Format("2006-01-02")
}

func Time2BJS(value time.Time) time.Time {
	return value.In(BJS)
}
