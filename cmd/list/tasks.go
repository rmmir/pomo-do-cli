package list

import (
	"fmt"

	m "github.com/rmmir/pomo-do/models"
	db "github.com/rmmir/pomo-do/database"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List all tasks",
	Long: `List all tasks in the task management system.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.ConnectDB()

		var tasks []m.Task
		result := db.DB.Find(&tasks)
		if result.Error != nil {
			fmt.Printf("Error fetching tasks: %v\n", result.Error)
			return
		}

		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("%v: %s\n", task.ID, task.Description)
		}
	},
}

func init() {}
