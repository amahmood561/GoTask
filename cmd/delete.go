package cmd

import (
    "fmt"
    "strconv"

    "github.com/spf13/cobra"
    "github.com/yourusername/gotask/task"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [task ID]",
    Short: "Delete a task",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Invalid task ID")
            return
        }
        err = task.DeleteTask(id)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println("Task deleted.")
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}
