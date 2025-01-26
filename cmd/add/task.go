package add

import (
	"fmt"
	"time"

	m "github.com/rmmir/pomo-do/models"
	db "github.com/rmmir/pomo-do/database"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var TaskCmd = &cobra.Command{
	Use:   "task [task description]",
	Short: "Add a new task",
	Long: `Add a new task to the list of tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.ConnectDB()
		if len(args) < 1 {
			fmt.Println("Please provide a task description.")
			return
		}

		task := m.Task{
			Description: args[0],
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Completed: false,
		}

		result := db.DB.Create(&task)
		if result.Error != nil {
			fmt.Printf("Error adding task: %v\n", result.Error)
			return
		}

		fmt.Printf("Task added with ID: %v - %s", task.ID, task.Description)
	},
}

func init() {}
