package fsm

import "time"

const (
	TimestampFormat = "2006-01-02T15:04:05.000000000-0700"
)

// TimestampToTime takes the timestamp represented as nanoseconds past epoch and converts it into a
// time.Time struct
func TimestampToTime(t int64) time.Time {
	return time.Unix(0, t)
}

// TimestampToString take the timestamp represented as nanoseconds past epoch and converts it to the
// a string using the Timestamp format
func TimestampToString(t int64) string {
	return TimestampToTime(t).Format(TimestampFormat)
}
