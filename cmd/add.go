package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "use this command to add tasks to your list of tasks",
	Long: `use this command to add a task through:

	togo add "do dishes" which will allow you to add any task you 
	desire but make sure to put any of your tasks inside of double quo-
	-tes ""`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0{
			fmt.Println("Please enter a task")
			return 
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
