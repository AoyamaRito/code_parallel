package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"make_parallel/internal/config"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Manage API configuration",
}

var apiSetCmd = &cobra.Command{
	Use:   "set [api-key]",
	Short: "Set Gemini API key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := args[0]
		if err := config.SetAPIKey(apiKey); err != nil {
			return fmt.Errorf("failed to set API key: %w", err)
		}
		fmt.Println("API key set successfully")
		return nil
	},
}

func init() {
	apiCmd.AddCommand(apiSetCmd)
}