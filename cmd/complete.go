package cmd

import (
    "fmt"
    "strconv"

    "github.com/spf13/cobra"
    "github.com/yourusername/gotask/task"
)

var completeCmd = &cobra.Command{
    Use:   "complete [task ID]",
    Short: "Mark a task as completed",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid task ID")
            return
        }
        err = task.CompleteTask(id)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Task marked as completed.")
    },
}

func init() {
    rootCmd.AddCommand(completeCmd)
}
