package list

import (
	"fmt"

	m "github.com/rmmir/pomo-do/models"
	db "github.com/rmmir/pomo-do/database"
	"github.com/spf13/cobra"
)

var categoryID int

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Edit a category to your task management system",
	Long: `The 'edit category' command allows you to add a category to your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()

		if categoryID > 0 {
			return listCategoryByID(categoryID)
		} else {
			return listAllCategories()
		}
	},
}

func init() {
	categoriesCmd.Flags().IntVar(&categoryID, "id", 0, "ID of the category to list")
}

func listAllCategories() error {
	var categories []m.Category
	result := db.DB.Find(&categories)
	if result.Error != nil {
		return fmt.Errorf("issues fetching categories: %v", result.Error)
	}

	fmt.Println("Categories:")
	for _, category := range categories {
		fmt.Printf("%v: %s\n", category.ID, category.Name)
	}

	return nil
}

func listCategoryByID(id int) error {
	var category m.Category
	result := db.DB.Find(&category, id)
	if result.Error != nil {
		return fmt.Errorf("issues fetching category with ID: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no category found with ID: %d", id)
	}

    fmt.Printf("%d: %s\n", category.ID, category.Name)
	return nil
}
