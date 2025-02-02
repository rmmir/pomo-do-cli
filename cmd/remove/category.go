package remove

import (
	"github.com/spf13/cobra"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Edit a category to your task management system",
	Long: `The 'edit category' command allows you to add a category to your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
}
