package utils

import "time"

func convertDate(date time.Time) string {
	return date.Format("January 2, 2006")
}

func getTime(date time.Time) string {
	return date.Format("15:04")
}

func GetThaiDate(date time.Time) string {
	return convertDate(date) + " เมื่อ " + getTime(date)
}
