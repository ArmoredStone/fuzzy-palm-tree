/*
Copyright © 2024 Alex Aheyev aheyevalex@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("Error reading file flag:", err)
			os.Exit(1)
		}
		fmt.Printf("Added todo %v to file: %v\n", args, file)
	},
}

// Here you define flags and configuration settings.
func init() {
	rootCmd.AddCommand(addCmd)
}
