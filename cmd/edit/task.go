package edit

import (
	"fmt"
	"time"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/spf13/cobra"
)

var taskID string

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Edit a task in your task management system",
	Long:  `The 'edit task' command allows you to edit a task in your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(taskID) != 8 {
			return fmt.Errorf("please provide a valid task ID with 8 characters")
		}

		if len(args) != 1 {
			return fmt.Errorf("please provide a new task description enclosed in quotes")
		}

		db.ConnectDB()
		newDescription := args[0]
		result := db.DB.Model(&m.Task{}).Where("id LIKE ?", "%"+taskID+"%").Updates(&m.Task{Description: newDescription, UpdatedAt: time.Now()})
		if result.Error != nil {
			return fmt.Errorf("issues updating the task - %v", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no task found with ID %s", taskID)
		}

		fmt.Printf("Task with ID: %s updated successfully to: %s\n", taskID, newDescription)
		return nil
	},
}

func init() {
	taskCmd.Flags().StringVar(&taskID, "id", "", "ID of the task to edit")
	taskCmd.MarkFlagRequired("id")
}
