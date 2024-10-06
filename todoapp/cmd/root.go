/*
Package cmd is the entry point for the application. It defines the root command and its subcommands.
*/

package cmd

import (
	"os"

	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todoapp",
	Short: "A brief description of your application",
	Long: `An application to manage your todo list.
	Using command line interface, you can add, remove,
	list and mark todo items as done.`,
	Run: func(cmd *cobra.Command, args []string) {
		toggle, err := cmd.Flags().GetBool("toggle")
		if err != nil {
			fmt.Println("Error reading toggle flag:", err)
			os.Exit(1)
		}
		fmt.Printf("Hello from rootCmd. Toggle is set to: %v\n", toggle)
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

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-cli.yaml)")
	rootCmd.PersistentFlags().BoolP("goggle", "g", false, "Help message for goggle")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
