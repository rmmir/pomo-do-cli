package edit

import (
	"fmt"

	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit tasks or categories in your task management system",
	Long:  `The 'edit' command allows you to edit tasks or categories in your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("'edit' command requires a subcommand (e.g., 'task' or 'category')")
	},
}

func init() {
	EditCmd.AddCommand(taskCmd)
}
