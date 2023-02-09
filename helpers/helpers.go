package helpers

import (
	"strconv"
	"strings"
	"time"
)

func ParseNowTime() string {
	now := time.Now()

	year := strconv.Itoa(now.Local().Year())
	month := now.Local().Month().String()
	day := strconv.Itoa(now.Local().Day())

	hour := strconv.Itoa(now.Local().Hour())
	min := strconv.Itoa(now.Local().Minute())
	sec := strconv.Itoa(now.Local().Second())

	clockTime := strings.Join([]string{hour, min, sec}, ":")
	date := strings.Join([]string{month, day, year}, " ")

	return strings.Join([]string{date, clockTime}, " ")
}
