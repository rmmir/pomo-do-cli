package view

import (
	"fmt"
	"log"

	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type TaskDisplay struct {
	ID          uint
	Description string
	Completed   bool
	Category    string
}

var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Interactive view of tasks",
	Long: `View tasks in an interactive mode. For example, you can select a task and then choose to edit or delete it.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var tasks []m.Task
		db.ConnectDB()
		tasksResult := db.DB.Preload("Category").Find(&tasks)
		if tasksResult.Error != nil {
			return fmt.Errorf("issues fetching tasks: %v", tasksResult.Error)
		}

		var tasksToDisplay []TaskDisplay
		for i, task := range tasks {
			tasksToDisplay = append(tasksToDisplay, TaskDisplay{
				ID:          uint(i),
				Description: task.Description,
				Completed:   task.Completed,
				Category:    task.Category.Name,
			})
		}

		prompt := promptui.Select{
			Label: "List of tasks (Select a task to edit or delete)",
			Items: tasksToDisplay,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "{{ if .Category }}\U0001F345 {{ .ID }}. {{ .Description | cyan }} ({{ .Category | cyan }}){{ else }}\U0001F345 {{ .ID }}. {{ .Description | cyan }} (No Category){{ end }}",
				Inactive: "{{ if .Category }}  {{ .ID }}. {{ .Description }} ({{ .Category }}){{ else }}  {{ .ID }}. {{ .Description }} (No Category){{ end }}",
				Selected: "{{ if .Category }}\U0001F345 {{ .ID }}. {{ .Description | red }} ({{ .Category | red }}){{ else }}\U0001F345 {{ .ID }}. {{ .Description | red }} (No Category){{ end }}",
			},
		}

		i, _, err := prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v\n", err)
		}

		selectedTask := tasksToDisplay[i]
		fmt.Printf("You selected: %s\n", selectedTask.Description)

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
			editTask(selectedTask)
		case "Delete":
			deleteTask(selectedTask)
		default:
			fmt.Println("Cancelled")
		}

		return nil
	},
}

func init() {
}

func editTask(task TaskDisplay) error {
	fmt.Println("Editing task:", task)

	prompt := promptui.Prompt{
        Label:   "Add new Description",
    }

    newDescription, err := prompt.Run()
    if err != nil {
        return fmt.Errorf("prompt failed: %v", err)
    }

	fmt.Println("New Description:", newDescription)
	return nil
}

func deleteTask(task TaskDisplay) error {
	fmt.Println("Deleting task:", task)

	return nil
}
