/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"time"

	"github.com/LombardiDaniel/go-time/parser"
	"github.com/LombardiDaniel/go-time/timer"
	"github.com/spf13/cobra"
)

var (
	font string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gotime [hh:mm:ss]",
	Args: cobra.MinimumNArgs(1),
	Short: "A terminal timer written in Go, it's Go-Time!",
	Long: `Go-Time is a terminal timer to help in anything such as flash talks and more!`,
	Example: `gotime 10 > inits 10 second
gotime 15:00 > inits 15 minute timer
gotime 01:15:00 > inits 1h15 minute timer
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var duration 	*time.Duration
		var err 		error

		if len(args) > 0 {
			timeStr := args[0]
			duration, err = parser.ParseTimeStr(timeStr)
			if err != nil {
				panic(err)
			}
		}

		timer.ShowTime(*duration, font)
    },
// 	Long: `A longer description that spans multiple lines and likely contains
// examples and usage of using your application. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-time.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// rootCmd.Flags().Int32("min", 0, "minutes")
	// rootCmd.Flags().Int32("sec", 0, "seconds")
	rootCmd.Flags().StringVarP(&font, "font", "f", "big", "Set the font, check: http://www.figlet.org/fontdb.cgi")
}


