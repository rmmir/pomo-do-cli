package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tasks or categories to your task management system",
	Long:  "The 'add' command allows you to add tasks or categories to your task management system. It requires a subcommand like 'task' or 'category'.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("'add' command requires a subcommand (e.g., 'task' or 'category')")
	},
}

func init() {
	AddCmd.AddCommand(taskCmd)
	AddCmd.AddCommand(categoryCmd)
}
