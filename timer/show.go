package timer

import (
	"fmt"
	"time"

	"github.com/LombardiDaniel/go-time/common"
	"github.com/LombardiDaniel/go-time/models"
	"github.com/common-nighthawk/go-figure"
)

const sleepBufferMs int = 100

type Opts struct {
	Font		string
}

func ShowTime(d time.Duration, font string) {
	tgtTime := time.Now().Add(d)

	var display models.TimeDisplay

	for tgtTime.After(time.Now()) {
		remainder := time.Until(tgtTime)
		display = models.NewTimeDisplay(remainder)		
		common.ClearScreen()

		switch remainderPart := float64(remainder.Milliseconds()) / float64(d.Milliseconds()); {
		case remainderPart <= 0.5 && remainderPart > 0.2:
			common.SetColor(common.CONSOLE_TEXT_COLOR_YELLOW)
		
		case remainderPart <= 0.20:
			common.SetColor(common.CONSOLE_TEXT_COLOR_RED)
		}

		str := fmt.Sprintf("%s:%s:%s", display.Hours, display.Minutes, display.Seconds)
		fig := figure.NewFigure(str, font, false)
		fig.Print()
		time.Sleep(time.Duration(sleepBufferMs) * time.Millisecond)
	}
}
