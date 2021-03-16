package utils

import "time"

func TimeInMilliseconds(t time.Time) int64 {
	return t.UnixNano() / 1000000
}
