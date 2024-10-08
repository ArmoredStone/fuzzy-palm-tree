/*
Copyright © 2024 Alex Aheyev aheyevalex@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve flags
		description, _ := cmd.Flags().GetString("desc")
		dueDate, _ := cmd.Flags().GetString("date")
		isDone, _ := cmd.Flags().GetBool("done")
		filePath, _ := cmd.Flags().GetString("file")

		// New TodoEnity
		task := NewTodoEntity(description, dueDate, isDone)
		var tasks []TodoEntity
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&tasks)
		if err != nil && err.Error() != "EOF" {
			fmt.Printf("Error decoding JSON: %v\n", err)
			os.Exit(1)
		}

		// Add the new task to the list
		tasks = append(tasks, *task)

		// Encode and write tasks to the file
		file.Seek(0, 0) // Reset the file pointer to the beginning
		encoder := json.NewEncoder(file)
		err = encoder.Encode(tasks)
		if err != nil {
			fmt.Printf("Error encoding JSON: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Added todo %v to file: %v\n", task.Description, filePath)
	},
}

// Here you define flags and configuration settings.
func init() {
	rootCmd.AddCommand(addCmd)
	description := ""
	dueDate := ""
	isDone := false
	addCmd.Flags().StringVar(&description, "desc", description, "Description of the task")
	addCmd.Flags().StringVar(&dueDate, "date", dueDate, "Due date of the task")
	addCmd.Flags().BoolVar(&isDone, "done", false, "Status of the task")
}
