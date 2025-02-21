package list

import (
	"fmt"

	m "github.com/rmmir/pomo-do/models"
	db "github.com/rmmir/pomo-do/database"
	u "github.com/rmmir/pomo-do/utils"
	"github.com/spf13/cobra"
)

var categoryID string

var categoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Edit a category to your task management system",
	Long: `The 'edit category' command allows you to add a category to your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()

		if len(categoryID) > 0 {
			return listCategoryByID(categoryID)
		} else {
			return listAllCategories()
		}
	},
}

func init() {
	categoriesCmd.Flags().StringVar(&categoryID, "id", "", "ID of the category to list")
}

func listAllCategories() error {
	var categories []m.Category
	result := db.DB.Find(&categories)
	if result.Error != nil {
		return fmt.Errorf("issues fetching categories: %v", result.Error)
	}

	if len(categories) == 0 {
		fmt.Println("No categories found")
		return nil
	}

	fmt.Println("Categories:")
	for _, category := range categories {
		fmt.Printf("%s: %s\n", u.GetShortUUID(category.ID), category.Name)
	}

	return nil
}

func listCategoryByID(id string) error {
	if len(id) != 8 {
		return fmt.Errorf("please provide a valid task ID with 8 characters")
	}

	var category m.Category
	result := db.DB.Find(&category, "id LIKE ?", "%"+id+"%")
	if result.Error != nil {
		return fmt.Errorf("issues fetching category with ID: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no category found with ID: %s", id)
	}

    fmt.Printf("%s: %s\n", id, category.Name)
	return nil
}
