/*
Copyright © 2024 Alex Aheyev aheyevalex@gmail.com
*/
package cmd

import (
	"bufio"
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
		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Printf("Error getting 'file' flag: %v\n", err)
			return
		}

		// Open the file if not opening try to create it
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			file, err = os.Create(filePath)
			if err != nil {
				fmt.Printf("Error creating file: %v\n", err)
				os.Exit(1)
			}
		}
		defer file.Close()

		// Print console messages
		fmt.Println("File:", filePath)

		// Read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
	},
}

// Here you define your flags and configuration settings.
func init() {
	rootCmd.AddCommand(listCmd)
}
