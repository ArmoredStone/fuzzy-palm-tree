/*
Package cmd is the entry point for the application. It defines the root command and its subcommands.

Copyright © 2024 Alex Aheyev aheyevalex@gmail.com
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// represents the file path to the text file with todo list items stored
var filePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todoapp",
	Short: "A brief description of your application",
	Long: `An application to manage your todo list.
	Using command line interface, you can add, remove,
	list and mark todo items as done.`,
	Run: func(cmd *cobra.Command, args []string) {
		listCmd.Run(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Here you define your flags and configuration settings.
func init() {
	rootCmd.PersistentFlags().StringVar(&filePath, "file", "mytodo.json", "Path to the file")
}
