package cmd

import (
	"fmt"
	"os"
	"encoding/json"
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

		tasks := make(map[string]bool)
		task := args[0]

		tasks[task] = false

		tasksJson, err := json.MarshalIndent(tasks, "", "  ")	
		if err != nil{
			fmt.Println("encountered an error while making the json file")
			return 
		}

		err = os.WriteFile("tasks.json", tasksJson, 0644)

		if err != nil{
			fmt.Println("encountered an error while writing the json file")
			return 
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
