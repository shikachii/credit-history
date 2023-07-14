package lib

import (
	"time"
)

type UnixTime int64

// Now returns the current time in Japan
func Now() time.Time {
	return time.Now().UTC().Add(time.Duration(9) * time.Hour)
}

// NowTimestamp returns the current time in Japan as a string
func NowTimestamp() string {
	return Now().Format("20060102150405")
}

// 実行した日の0:00のUNIX時間を返す
func Today() int64 {
	now := Now()
	utime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()
	return utime
}

// UnixTime型(int64)を文字列に変換する
func (ut UnixTime) String() string {
	return time.Unix(int64(ut), 0).Format("20060102150405")
}
