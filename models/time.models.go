package models

import (
	"strconv"
	"time"
)


type TimeDisplay struct {
	Hours				string
	Minutes				string
	Seconds				string
}

func NewTimeDisplay(d time.Duration) TimeDisplay {
	return TimeDisplay{
		Hours: formatNum(d.Hours(), 60),
		Minutes: formatNum(d.Minutes(), 60),
		Seconds: formatNum(d.Seconds(), 60),
	}
}

func formatNum(n float64, max int32) string {
	if n < 1 {
		return "00"
	}

	mod := int(n) % int(max)

	return strconv.Itoa(mod)
}