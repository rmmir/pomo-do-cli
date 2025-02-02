package add

import (
	"github.com/spf13/cobra"
)

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Add a category to your task management system",
	Long: `The 'add category' command allows you to add a category to your task management system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
}
