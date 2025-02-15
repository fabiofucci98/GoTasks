package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GoTasks",
	Short: "A simple task manager CLI",
	Long:  "GoTasks is a command-line application to add, list, and delete tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Task Manager CLI! Use --help to see available commands.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
