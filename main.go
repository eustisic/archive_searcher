package main

import (
	"fmt"
	"milestones/commands"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archive-cli",
	Short: "archive-cli is a tool for searching the Internet Archive",
	Run: func(cmd *cobra.Command, args []string) {

		commands.ListRecordings()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
