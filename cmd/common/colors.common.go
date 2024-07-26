package common

import (
	"fmt"
)


const (
	CONSOLE_TEXT_COLOR_BLACK  	string = "\033[30m"
	CONSOLE_TEXT_COLOR_RED    	string = "\033[31m"
	CONSOLE_TEXT_COLOR_GREEN  	string = "\033[32m"
	CONSOLE_TEXT_COLOR_YELLOW 	string = "\033[33m"
	CONSOLE_TEXT_COLOR_BLUE   	string = "\033[34m"
	CONSOLE_TEXT_COLOR_PURPLE 	string = "\033[35m"
	CONSOLE_TEXT_COLOR_CYAN   	string = "\033[36m"
	CONSOLE_TEXT_COLOR_WHITE  	string = "\033[37m"

	CONSOLE_COLOR_RESET			string = "\033[0m"
)

func PrintOnColor(t string, c string) error {
	_, err := fmt.Printf("%s%s%s", c, t, CONSOLE_COLOR_RESET)
	return err
}

func SetColor(c string) error {
	_, err := fmt.Print(c)
	return err
}

func ResetColor() error {
	_, err := fmt.Print(CONSOLE_COLOR_RESET)
	return err
}