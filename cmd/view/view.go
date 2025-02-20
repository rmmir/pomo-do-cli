package view

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	db "github.com/rmmir/pomo-do/database"
	m "github.com/rmmir/pomo-do/models"
	"github.com/spf13/cobra"
)

type TaskDisplay struct {
	ID          uuid.UUID
	Position    uint
	Description string
	Completed   bool
	Category    string
}

var ViewCmd = &cobra.Command{
	Use:   "view",
	Short: "Interactive view of tasks",
	Long:  `View tasks in an interactive mode. For example, you can select a task and then choose to edit or delete it.`,
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
				ID:          task.ID,
				Position:    uint(i+1),
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
				Active:   "{{ if .Category }}\U0001F345 {{ .Position }}. {{ .Description | cyan }} ({{ .Category | cyan }}){{ else }}\U0001F345 {{ .Position }}. {{ .Description | cyan }} (No Category){{ end }}",
				Inactive: "{{ if .Category }}  {{ .Position }}. {{ .Description }} ({{ .Category }}){{ else }}  {{ .Position }}. {{ .Description }} (No Category){{ end }}",
				Selected: "{{ if .Category }}\U0001F345 {{ .Position }}. {{ .Description | red }} ({{ .Category | red }}){{ else }}\U0001F345 {{ .Position }}. {{ .Description | red }} (No Category){{ end }}",
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

func init() {}

func editTask(task TaskDisplay) error {
	prompt := promptui.Prompt{
		Label: "Add new Description",
	}

	newDescription, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %v", err)
	}

	result := db.DB.Model(&m.Task{}).Where("ID = ?", task.ID).Updates(&m.Task{Description: newDescription, UpdatedAt: time.Now()})
	if result.Error != nil {
		return fmt.Errorf("issues updating the task - %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no task found with ID %s", task.ID)
	}

	fmt.Println("New Description:", newDescription)
	return nil
}

func deleteTask(task TaskDisplay) error {
	prompt := promptui.Select{
		Label: "Are you sure you want to delete this task? (yes/no)",
		Items: []string{"yes", "no"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %v", err)
	}

	if result == "no" {
		fmt.Println("Task deletion was cancelled")
	}

	if result == "yes" {
		result := db.DB.Delete(&m.Task{}, task.ID)
		if result.Error != nil {
			return fmt.Errorf("issues deleting the task - %v", result.Error)
		}

		if result.RowsAffected == 0 {
            return fmt.Errorf("no task found with ID %s", task.ID)
        }

        fmt.Println("Task deleted successfully")
	}

	return nil
}
