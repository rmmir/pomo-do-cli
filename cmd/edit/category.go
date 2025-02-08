package edit

import (
	"fmt"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/spf13/cobra"
)

var categoryID string

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Edit a category to your task management system",
	Long:  `The 'edit category' command allows you to add a category to your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(categoryID) != 8 {
			return fmt.Errorf("please provide a valid task ID with 8 characters")
		}
		
		if len(args) != 1 {
			return fmt.Errorf("please provide a new category description enclosed in quotes")
		}
		
		db.ConnectDB()

		newCategoryName := args[0]
		result := db.DB.Model(&m.Category{}).Where("id LIKE ?", "%"+categoryID+"%").Update("name", newCategoryName)
		if result.Error != nil {
			return fmt.Errorf("issues updating the category - %v", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no category found with ID %s", categoryID)
		}

		fmt.Printf("Task with ID: %s updated successfully to: %s\n", categoryID, newCategoryName)
		return nil
	},
}

func init() {
	categoryCmd.Flags().StringVar(&categoryID, "id", "", "ID of the category to edit")
	categoryCmd.MarkFlagRequired("id")
}
