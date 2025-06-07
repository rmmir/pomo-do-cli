package add

import (
	"fmt"
	"time"

	m "github.com/rmmir/pomo-do/models"
	db "github.com/rmmir/pomo-do/database"
	"github.com/spf13/cobra"
)

var workTime time.Duration
var breakTime time.Duration
var repetitions int

var sessionCmd = &cobra.Command{
	Use:   "task [task description]",
	Short: "Add a new task",
	Long: `Add a new task to the list of tasks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db.ConnectDB()
		if len(args) < 1 {
			return fmt.Errorf("please provide a task description")
		}

		if len(args) != 1 {
			return fmt.Errorf("please provide a new task description enclosed in quotes")
		}

		_ = m.Session{
			Name: args[0],
			WorkTime:    workTime,
			BreakTime:   breakTime,
			Repetitions: repetitions,
		}

		return fmt.Errorf("session command is not implemented yet")
	},
}

func init() {
	sessionCmd.Flags().DurationVar(&workTime, "workTime", 45*time.Minute, "Work duration for the session")
	sessionCmd.Flags().DurationVar(&breakTime, "breakTime", 15*time.Minute, "Break duration for the session")
	sessionCmd.Flags().IntVar(&repetitions, "repetitions", 1, "Number of repetitions for the session")
}