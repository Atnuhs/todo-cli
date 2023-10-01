package internal

import "time"

func Basename(t time.Time) string {
	return t.Format("20060102")
}
