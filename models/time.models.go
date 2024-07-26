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
		Hours: FormatNum(d.Hours(), 60),
		Minutes: FormatNum(d.Minutes(), 60),
		Seconds: FormatNum(d.Seconds(), 60),
	}
}

func FormatNum(n float64, max int32) string {
	if n < 1 {
		return "00"
	}

	mod := int(n) % int(max)

	val := strconv.Itoa(mod)
	if mod > 9 {
		return val
	}

	return "0" + val
}