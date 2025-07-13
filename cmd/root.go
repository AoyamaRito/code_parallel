package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "make_parallel",
	Short: "Parallel AI task execution tool",
	Long:  "A tool for executing AI tasks in parallel using Gemini API",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(queueCmd)
}