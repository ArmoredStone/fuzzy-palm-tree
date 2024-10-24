/*
Copyright © 2024 Alex Aheyev aheyevalex@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists created TODO items",
	Long:  `A list of created TODO items`,
	Run: func(cmd *cobra.Command, args []string) {
		// getting filePath flag
		filePath, _ := cmd.Flags().GetString("file")

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		// Read and decode tasks from the file
		var tasks []TodoEntity
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&tasks)
		if err != nil {
			fmt.Printf("Error decoding JSON: %v\n", err)
			os.Exit(1)
		}

		// Print the tasks
		for _, task := range tasks {
			fmt.Printf("Name: %s (Due: %s, Done: %v)\n", task.Name, task.DueDate, task.IsDone)
		}
	},
}

// Here you define your flags and configuration settings.
func init() {
	rootCmd.AddCommand(listCmd)
}
