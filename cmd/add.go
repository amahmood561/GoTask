package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/yourusername/gotask/task"
)

var addCmd = &cobra.Command{
    Use:   "add [task description]",
    Short: "Add a new task",
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        description := args[0]
        newTask := task.Task{
            ID:          task.GetNextID(),
            Description: description,
            Completed:   false,
        }
        task.AddTask(newTask)
        fmt.Println("Added task:", description)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
