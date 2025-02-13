package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet is a simple CLI tool",
	Long:  "This is a command-line tool that greets the user.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

var nameCmd = &cobra.Command{
	Use:   "name [name]",
	Short: "Print a custom greeting",
	Long:  "Prints a greeting with the name passed as argument",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello, %s!\n", args[0])
	},
}

func main() {
	rootCmd.AddCommand(nameCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
