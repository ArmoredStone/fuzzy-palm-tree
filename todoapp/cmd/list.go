/*
Copyright © 2024 Alex Aheyev aheyevalex@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists created TODO items",
	Long:  `A list of created TODO items`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("Error getting 'file' flag:", err)
			return
		}
		fmt.Println("File:", file)
		fmt.Println("List todo called")
	},
}

// Here you define your flags and configuration settings.
func init() {
	rootCmd.AddCommand(listCmd)
}
