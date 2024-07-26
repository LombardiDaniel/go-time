package timer

import (
	"fmt"
	"time"

	"github.com/LombardiDaniel/go-time/cmd/common"
	"github.com/LombardiDaniel/go-time/cmd/models"
	"github.com/common-nighthawk/go-figure"
)

func ShowTime(d time.Duration) {
	tgtTime := time.Now().Add(d)

	var display models.TimeDisplay

	for tgtTime.After(time.Now()) {
		common.ClearScreen()
		remainder := time.Until(tgtTime)
		display = models.NewTimeDisplay(remainder)
		
		switch remainderPart := float64(remainder.Milliseconds()) / float64(d.Milliseconds()); {
		case remainderPart <= 0.5 && remainderPart > 0.2:
			common.SetColor(common.CONSOLE_TEXT_COLOR_YELLOW)
		
		case remainderPart <= 0.20:
			common.SetColor(common.CONSOLE_TEXT_COLOR_RED)
		}


		str := fmt.Sprintf("%s:%s:%s", display.Hours, display.Minutes, display.Seconds)
		fig := figure.NewFigure(str, "", false)
		fig.Print()
		time.Sleep(100 * time.Millisecond)
	}
}
