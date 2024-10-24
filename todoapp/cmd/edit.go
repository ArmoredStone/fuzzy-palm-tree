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
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edits by name",
	Long:  `Edit desired task by accessing it by name.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve flags
		name, _ := cmd.Flags().GetString("name")
		dueDate, _ := cmd.Flags().GetString("date")
		dueDateValue := cmd.Flags().Changed("date")
		isDone, _ := cmd.Flags().GetBool("done")
		isDoneValue := cmd.Flags().Changed("done")
		filePath, _ := cmd.Flags().GetString("file")

		editedEntity := TodoEntity{name, dueDate, isDone}

		// New TodoEnity
		var tasks []TodoEntity
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&tasks)
		if err != nil && err.Error() != "EOF" {
			fmt.Printf("Error decoding JSON: %v\n", err)
			os.Exit(1)
		}

		// Loop to find the task by name
		// and apply the changes
		for i := range tasks {
			if tasks[i].Name == editedEntity.Name {
				buf := tasks[i]
				if dueDateValue {
					tasks[i].DueDate = editedEntity.DueDate
				}
				if isDoneValue {
					tasks[i].IsDone = editedEntity.IsDone
				}
				fmt.Printf("Changed %v\n into %v\n", buf, tasks[i])
			}
		}

		// Encode and write tasks to the file
		file.Truncate(0)
		file.Seek(0, 0) // Reset the file pointer to the beginning
		encoder := json.NewEncoder(file)
		err = encoder.Encode(tasks)
		if err != nil {
			fmt.Printf("Error encoding JSON: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

	},
}

// Here you define flags and configuration settings.
func init() {
	rootCmd.AddCommand(editCmd)
	name := ""
	dueDate := ""
	isDone := false
	editCmd.Flags().StringVar(&name, "name", name, "Name of the task")
	editCmd.Flags().StringVar(&dueDate, "date", dueDate, "Due date of the task")
	editCmd.Flags().BoolVar(&isDone, "done", isDone, "Status of the task")
}
