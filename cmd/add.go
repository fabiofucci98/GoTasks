package cmd

import (
	s "GoTasks/structs"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func addTask(description, dueDate, difficulty string) {
	task := s.Task{
		Description: description,
		DueDate:     dueDate,
		Difficulty:  difficulty,
	}

	filePath := "tasks.json"
	var tasks []s.Task

	if _, err := os.Stat(filePath); err == nil {
		file, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		err = json.Unmarshal(file, &tasks)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}
	}

	tasks = append(tasks, task)

	file, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

var addCmd = &cobra.Command{
	Use:   "add [description] [dueDate] [difficulty]",
	Short: "Add a new task",
	Long:  "Adds a new task with a description, due date, and difficulty",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		addTask(args[0], args[1], args[2])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
