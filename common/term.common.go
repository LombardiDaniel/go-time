package common

import (
	"os"
	"os/exec"
	"strings"
)

var xTermSys bool = strings.Contains(os.Getenv("TERM"), "xterm")

func ClearScreen() {
	var cmd *exec.Cmd

	switch xTermSys {
	case true:
		cmd = exec.Command("clear")
	default:
		cmd = exec.Command("cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}