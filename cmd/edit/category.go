package edit

import (
	"fmt"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/spf13/cobra"
)

var categoryID int

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Edit a category to your task management system",
	Long:  `The 'edit category' command allows you to add a category to your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetInt("id")
		if err != nil || id == 0 {
			return fmt.Errorf("please provide a valid task ID: %v", err)
		}
		
		if len(args) != 1 {
			return fmt.Errorf("please provide a new category description enclosed in quotes")
		}
		
		db.ConnectDB()

		result := db.DB.Model(&m.Category{}).Where("id = ?", id).Update("name", args[0])
		if result.Error != nil {
			return fmt.Errorf("issues updating the category - %v", result.Error)
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no category found with ID %d", id)
		}

		fmt.Printf("Task with ID: %d updated successfully to: %s\n", categoryID, args[0])
		return nil
	},
}

func init() {
	categoryCmd.Flags().IntVar(&categoryID, "id", 0, "ID of the category to edit")
}
