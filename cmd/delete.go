package cmd

import (
	s "GoTasks/structs"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func deleteTask(description string) {
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

	for i, task := range tasks {
		if task.Description == description {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

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

var deleteCmd = &cobra.Command{
	Use:   "delete [description]",
	Short: "Delete a task",
	Long:  "Deletes a task with the given description",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		deleteTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
