package utils

import "time"

func DaysRemaining(date time.Time) int {
	return int(time.Until(date.Local()).Hours() / 24)
}

func DateFormated(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
