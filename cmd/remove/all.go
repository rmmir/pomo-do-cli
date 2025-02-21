package remove

import (
	"fmt"

	db "github.com/rmmir/pomo-do/database"
	"github.com/spf13/cobra"
)

var allTasksCmd = &cobra.Command{
	Use:   "all",
	Short: "Removes all tasks from your task management system",
	Long: `The 'remove all' command allows you to remove all tasks from your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()
		result := db.DB.Exec("DELETE FROM tasks")
		if result.Error != nil {
			return fmt.Errorf("issues removing all tasks - %v", result.Error)
		}

		fmt.Println("All tasks removed successfully")
		return nil
	},
}

var allCategoriesCmd = &cobra.Command{
	Use:   "all",
	Short: "Removes all categories from your task management system",
	Long: `The 'remove all' command allows you to remove all categories from your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()
		result := db.DB.Exec("DELETE FROM categories")
		if result.Error != nil {
			return fmt.Errorf("issues removing all categories - %v", result.Error)
		}

		fmt.Println("All categories removed successfully")
		return nil
	},
}

func init() {
	taskCmd.AddCommand(allTasksCmd)
	categoryCmd.AddCommand(allCategoriesCmd)
}