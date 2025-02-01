package remove

import (
	"fmt"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/spf13/cobra"
)

var taskID int

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Removes tasks or categories in your task management system",
	Long: `The 'remove task' command allows you to remove tasks or categories in your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id == 0 {
			return fmt.Errorf("please provide a valid task ID: %v", err)
		}

		db.ConnectDB()
		result := db.DB.Delete(&m.Task{}, taskID)
		if result.Error != nil {
			return fmt.Errorf("issues removing the task - %v", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no task found with ID %d", id)
		}

		fmt.Printf("Task with ID: %d removed successfully\n", id)

		return nil
	},
}

func init() {
	taskCmd.Flags().IntVar(&taskID, "id", 0, "ID of the task to remove")
}
