package utils

import "time"

const (
	DatetimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
)

// 字符串转时间
func GetTime(t string, format string) (r time.Time, err error) {
	return time.ParseInLocation(format, t, time.Local)
}
