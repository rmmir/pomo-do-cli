package remove

import (
	"fmt"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/spf13/cobra"
)

var categoryID string

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Remove a category from your task management system",
	Long: `The 'remove category' command allows you to remove a category from your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(categoryID) != 8 {
			return fmt.Errorf("please provide a valid task ID with 8 characters")
		}

		db.ConnectDB()
		result := db.DB.Delete(&m.Category{}, "id LIKE ?", "%"+categoryID+"%")
		if result.Error != nil {
			return fmt.Errorf("issues removing the category - %v", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no category found with ID %s", categoryID)
		}

		fmt.Printf("Category with ID: %s removed successfully\n", categoryID)

		return nil
	},
}

func init() {
	categoryCmd.Flags().StringVar(&categoryID, "id", "", "ID of the category to remove")
	categoryCmd.MarkFlagRequired("id")
}
