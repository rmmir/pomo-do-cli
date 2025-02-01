package list

import (
	"fmt"

	m "github.com/rmmir/pomo-do/models"
	db "github.com/rmmir/pomo-do/database"
	"github.com/spf13/cobra"
)

var taskID int

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List all tasks",
	Long: `List all tasks in the task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()

		if taskID > 0 {
			return listTaskByID(taskID)
		} else {
			return listAllTasks()
		}
	},
}

func init() {
	tasksCmd.Flags().IntVar(&taskID, "id", 0, "ID of the task to list")
}

func listAllTasks() error {
	var tasks []m.Task
	result := db.DB.Find(&tasks)
	if result.Error != nil {
		return fmt.Errorf("issues fetching tasks: %v", result.Error)
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("%v: %s\n", task.ID, task.Description)
	}

	return nil
}

func listTaskByID(id int) error {
	var task m.Task
	result := db.DB.Find(&task, id)
	if result.Error != nil {
		return fmt.Errorf("issues fetching task with ID: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no task found with ID: %d", id)
	}

    fmt.Printf("%d: %s\n", task.ID, task.Description)
	return nil
}
