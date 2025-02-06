package remove

import (
	"fmt"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/spf13/cobra"
)

var taskID string

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Removes tasks or categories in your task management system",
	Long: `The 'remove task' command allows you to remove tasks or categories in your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(taskID) != 8 {
			return fmt.Errorf("please provide a valid task ID with 8 characters")
		}
		
		db.ConnectDB()
		result := db.DB.Delete(&m.Task{}, "id LIKE ?", "%"+taskID+"%")
		if result.Error != nil {
			return fmt.Errorf("issues removing the task - %v", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no task found with ID %s", taskID)
		}

		fmt.Printf("Task with ID: %s removed successfully\n", taskID)

		return nil
	},
}

func init() {
	taskCmd.Flags().StringVar(&taskID, "id", "", "ID of the task to edit")
    taskCmd.MarkFlagRequired("id")
}
