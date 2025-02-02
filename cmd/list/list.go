package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items",
	Long: `List all items in the task management system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	ListCmd.AddCommand(tasksCmd)
	ListCmd.AddCommand(categoriesCmd)
}
