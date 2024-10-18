package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "gotask",
    Short: "GoTask is a simple CLI task manager",
    Long:  `A command-line task manager built with Go, allowing you to add, list, complete, and delete tasks.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    // Initialize configuration, if needed
}
