package add

import (
	"fmt"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	u "github.com/rmmir/pomo-do/utils"
	"github.com/spf13/cobra"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Add a category to your task management system",
	Long: `The 'add category' command allows you to add a category to your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()
		if len(args) < 1 {
			return fmt.Errorf("please provide a new category description")
		}

		if len(args) != 1 {
			return fmt.Errorf("please provide a new category description enclosed in quotes")
		}

		category := m.Category{
			Name: args[0],
		}
		
		result := db.DB.Create(&category)
		if result.Error != nil {
			return fmt.Errorf("unexpected issues when adding the category - %v", result.Error)
		}

		fmt.Printf("Category with ID: %s added successfully\n", u.GetShortUUID(category.ID))
		return nil
	},
}

func init() {
}
