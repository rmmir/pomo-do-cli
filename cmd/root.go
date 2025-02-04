package cmd

import (
	"os"

	db "github.com/rmmir/pomo-do/database"
	"github.com/rmmir/pomo-do/cmd/add"
	"github.com/rmmir/pomo-do/cmd/list"
	"github.com/rmmir/pomo-do/cmd/edit"
	"github.com/rmmir/pomo-do/cmd/remove"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pomo-do",
	Short: "Pomo-do is a CLI tool for managing your tasks using the Pomodoro Technique",
	Long: `Pomo-do is a CLI tool for managing your tasks using the Pomodoro Technique.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	db.ConnectDB()	
	addSubcommands()
}

func addSubcommands() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(edit.EditCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
}