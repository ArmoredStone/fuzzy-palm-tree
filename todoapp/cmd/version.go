/*
Copyright © 2024 Alex Aheyev aheyevalex@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "0.1.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A display of the current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

// Here you define your flags and configuration settings.
func init() {
	rootCmd.AddCommand(versionCmd)
}
