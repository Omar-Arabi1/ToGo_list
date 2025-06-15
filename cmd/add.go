package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
		
		jsonToTasks, errReading := os.ReadFile("tasks.json")

		if errReading != nil{
			fmt.Println("encountered an error while reading the json file")
			return 
		}
		
		tasks := make(map[string]bool)

		task := args[0]

		if len(strings.TrimSpace(task)) == 0{
			fmt.Println("Please enter a task")
			return 
		}

		tasks[task] = false

		errReading = json.Unmarshal(jsonToTasks, &tasks)
		if errReading != nil{
			fmt.Println("encountered an error while reading the json file")
			return 
		}

		tasksToJson, err := json.MarshalIndent(tasks, "", "  ")	

		if err != nil{
			fmt.Println("encountered an error while making the json file")
			return 
		}

		err = os.WriteFile("tasks.json", tasksToJson, 0644)

		if err != nil{
			fmt.Println("encountered an error while writing the json file")
			return 
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
