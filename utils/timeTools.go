package utils

import (
    "fmt"
    "time"
)

// Set time zone at UTC/GMT+08:00,  China Standard Time UT+8:00
var cnTimeZone = time.FixedZone("CST", 8*3600)

// Get the timestamp of now, unit: sec
func NowTimestamp() int64 {
    return time.Now().Unix()
}

// Get the datetime string of one timestamp
func DatetimeStrOfTimestamp(timestamp int64) string {
    return time.Unix(timestamp, 0).In(cnTimeZone).Format("2006-01-02 15:04:05")
}

// Get the datetime string of now
func NowDatetimeStr() string {
    return DatetimeStrOfTimestamp(NowTimestamp())
}

// Get the date string of today
func TodayDateStr() string {
    nowTime := time.Now().In(cnTimeZone)
    return fmt.Sprintf("%d-%02d-%02d", nowTime.Year(), nowTime.Month(), nowTime.Day())
}
