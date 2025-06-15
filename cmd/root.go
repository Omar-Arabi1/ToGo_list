package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Togo_list",
	Short: "this tool allows you to add, list, mark done and remove tasks",
	Long: `this is a cli tool that will allow you to list remove mark done and add your 
	tasks to a json file to stay organized!

	we hope you enjoy`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
		rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


