package common

import (
	"os"
	"os/exec"
	"strings"
)

func ClearScreen() {
	var cmd *exec.Cmd

	switch strings.Contains(os.Getenv("TERM"), "xterm") {
	case true:
		cmd = exec.Command("clear")
	default:
		cmd = exec.Command("cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}