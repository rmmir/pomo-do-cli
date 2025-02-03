package remove

import (
	"fmt"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/spf13/cobra"
)

var categoryID int

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Remove a category from your task management system",
	Long: `The 'remove category' command allows you to remove a category from your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id == 0 {
			return fmt.Errorf("please provide a valid task ID: %v", err)
		}

		db.ConnectDB()
		result := db.DB.Delete(&m.Category{}, id)
		if result.Error != nil {
			return fmt.Errorf("issues removing the category - %v", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no category found with ID %d", id)
		}

		fmt.Printf("Category with ID: %d removed successfully\n", id)

		return nil
	},
}

func init() {
	categoryCmd.Flags().IntVar(&categoryID, "id", 0, "ID of the category to remove")
}
