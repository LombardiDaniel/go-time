package parser

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/LombardiDaniel/go-time/cmd/common"
)

var (
	// returns 6 matches
	parseTimeReg *regexp.Regexp = regexp.MustCompilePOSIX(`^(([0-1][0-9]|2[0-3]):)?(([0-5][0-9]):)?([0-5][0-9])$`)
	ErrNotParse error = errors.New("could not parse datetimeStr")
)

const (
	matchFirstPosIdx 		int = 2
	matchSecondPosIdx 		int = 4
	matchThirdPostIdx		int = 5
)

func ParseTimeStr(timeStr string) (*time.Duration, error) {

	matches := parseTimeReg.FindAllStringSubmatch(timeStr, -1)
	if len(matches) == 0 {
		return nil, ErrNotParse
	}

	var hours, minutes, seconds int32
	emptiesCount := common.SliceContains(matches[0], "")

	switch emptiesCount {
	case 0:
		h, err := strconv.Atoi(matches[0][matchFirstPosIdx])
		if err != nil {
			return nil, err
		}
		hours = int32(h)

		m, err := strconv.Atoi(matches[0][matchSecondPosIdx])
		if err != nil {
			return nil, err
		}
		minutes = int32(m)

		s, err := strconv.Atoi(matches[0][matchThirdPostIdx])
		if err != nil {
			return nil, err
		}
		seconds = int32(s)

	case 2:
		hours = 0

		m, err := strconv.Atoi(matches[0][matchFirstPosIdx])
		if err != nil {
			return nil, err
		}
		minutes = int32(m)

		s, err := strconv.Atoi(matches[0][matchThirdPostIdx])
		if err != nil {
			return nil, err
		}
		seconds = int32(s)
	
	case 4:
		hours = 0
		minutes = 0

		s, err := strconv.Atoi(matches[0][matchThirdPostIdx])
		if err != nil {
			return nil, err
		}
		seconds = int32(s)

	default:
		return nil, ErrNotParse
	}

	t := time.Duration(hours) * time.Hour + time.Duration(minutes) * time.Minute + time.Duration(seconds) * time.Second
	return &t, nil
}