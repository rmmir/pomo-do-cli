package add

import (
	"fmt"
	"time"

	m "github.com/rmmir/pomo-do/models"
	db "github.com/rmmir/pomo-do/database"
	u "github.com/rmmir/pomo-do/utils"
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task [task description]",
	Short: "Add a new task",
	Long: `Add a new task to the list of tasks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()
		if len(args) < 1 {
			return fmt.Errorf("please provide a task description")
		}

		if len(args) != 1 {
			return fmt.Errorf("please provide a new task description enclosed in quotes")
		}

		task := m.Task{
			Description: args[0],
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Completed: false,
		}

		result := db.DB.Create(&task)
		if result.Error != nil {
			return fmt.Errorf("unexpected issues when adding the task - %v", result.Error)
		}

		fmt.Printf("Task added with ID: %v - %s\n", u.GetShortUUID(task.ID), task.Description)
		return nil
	},
}

func init() {}
