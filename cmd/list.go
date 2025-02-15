package cmd

import (
	s "GoTasks/structs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func listTasks() {
	// Open the tasks.json file
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatalf("Failed to open tasks.json: %v", err)
	}
	defer file.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read tasks.json: %v", err)
	}

	// Parse the JSON data
	var tasks []s.Task
	if err := json.Unmarshal(byteValue, &tasks); err != nil {
		log.Fatalf("Failed to parse tasks.json: %v", err)
	}

	// Print the tasks
	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Description)
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks stored in the task manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
