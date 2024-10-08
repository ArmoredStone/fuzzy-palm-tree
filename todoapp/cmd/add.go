/*
Copyright © 2024 Alex Aheyev aheyevalex@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	//"encoding/json"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		//filepath flag retrieval
		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Printf("Error reading file flag: %v\n", err)
			os.Exit(1)
		}

		// Open the file if not opening try to create it
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			file, err = os.Create(filePath)
			if err != nil {
				fmt.Printf("Error creating file: %v\n", err)
				os.Exit(1)
			}
		}
		defer file.Close()

		// join the arguments with comma
		task := strings.Join(args, ", ")

		// Write the joined arguments to the file
		if _, err := file.WriteString(task + "\n"); err != nil {
			fmt.Printf("could not write to file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Added todo %v to file: %v\n", task, filePath)
	},
}

// Here you define flags and configuration settings.
func init() {
	rootCmd.AddCommand(addCmd)
	//	addCmd.Flags().AddFlag().String("file", "todo.txt", "File to store TODO items")

}
