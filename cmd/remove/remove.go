package remove

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove tasks or categories from your task management system",	
	Long: `The 'remove' command allows you to remove tasks or categories from your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("'remove' command requires a subcommand (e.g., 'task' or 'category')")
	},
}

func init() {
	RemoveCmd.AddCommand(taskCmd)
	RemoveCmd.AddCommand(categoryCmd)
}
