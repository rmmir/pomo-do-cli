package view

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
)

var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Interactive view of tasks",
	Long: `View tasks in an interactive mode. For example, you can select a task and then choose to edit or delete it.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var tasks []m.Task
		db.ConnectDB()
		tasksResult := db.DB.Find(&tasks)
		if tasksResult.Error != nil {
			return fmt.Errorf("issues fetching tasks: %v", tasksResult.Error)
		}

		var taskDescriptions []string
		for i, task := range tasks {
			taskDescriptions = append(taskDescriptions, fmt.Sprintf("%d. %s", i+1, task.Description))
		}

		prompt := promptui.Select{
			Label: "List of tasks (Select a task to edit or delete)",
			Items: taskDescriptions,
		}

		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v\n", err)
		}

		fmt.Printf("You selected: %s\n", result)

		actionPrompt := promptui.Select{
			Label: "What do you want to do?",
			Items: []string{"Edit", "Delete", "Cancel"},
		}

		_, action, err := actionPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v\n", err)
		}

		switch action {
		case "Edit":
			fmt.Println("Editing task:", result)
			// TODO: Implement edit logic
		case "Delete":
			fmt.Println("Deleting task:", result)
			// TODO: Implement delete logic
		default:
			fmt.Println("Cancelled")
		}

		return nil
	},
}

func init() {
}
