package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items",
	Long: `List all items in the task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("'list' command requires a subcommand (e.g., 'tasks' or 'categories')")
	},
}

func init() {
	ListCmd.AddCommand(tasksCmd)
	ListCmd.AddCommand(categoriesCmd)
}
