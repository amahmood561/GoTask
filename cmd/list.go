package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/yourusername/gotask/task"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        tasks := task.GetAllTasks()
        for _, t := range tasks {
            status := " "
            if t.Completed {
                status = "X"
            }
            fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Description)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
