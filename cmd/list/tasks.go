package list

import (
	"fmt"

	m "github.com/rmmir/pomo-do/models"
	db "github.com/rmmir/pomo-do/database"
	u "github.com/rmmir/pomo-do/utils"
	"github.com/spf13/cobra"
)

var taskID string

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List all tasks",
	Long: `List all tasks in the task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()
		
		if len(taskID) > 0 {
			return listTaskByID(taskID)
		} else {
			return listAllTasks()
		}
	},
}

func init() {
	tasksCmd.Flags().StringVar(&taskID, "id", "", "ID of the task to edit")
}

func listAllTasks() error {
	var tasks []m.Task
	result := db.DB.Preload("Category").Find(&tasks)
	if result.Error != nil {
		return fmt.Errorf("issues fetching tasks: %v", result.Error)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	fmt.Println("Tasks:")
	for _, task := range tasks {
		fmt.Printf("%v: %s\n", u.GetShortUUID(task.ID), task.Description)
	}

	return nil
}

func listTaskByID(id string) error {
	if len(id) != 8 {
		return fmt.Errorf("please provide a valid task ID with 8 characters")
	}

	var task m.Task
	result := db.DB.Preload("Category").Find(&task, "id LIKE ?", "%"+id+"%")
	if result.Error != nil {
		return fmt.Errorf("issues fetching task with ID: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no task found with ID: %s", id)
	}

    fmt.Printf("%s: %s\n", id, task.Description)
	return nil
}
