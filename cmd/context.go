package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"code_parallel/internal/config"
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Manage project context",
}

var contextSetCmd = &cobra.Command{
	Use:   "set [context-text]",
	Short: "Set project context",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		contextText := args[0]
		if err := config.SetContext(contextText); err != nil {
			return fmt.Errorf("failed to set context: %w", err)
		}
		fmt.Printf("プロジェクトコンテキストを設定しました: %s\n", contextText)
		return nil
	},
}

var contextShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current project context",
	RunE: func(cmd *cobra.Command, args []string) error {
		context, err := config.GetContext()
		if err != nil {
			return fmt.Errorf("failed to get context: %w", err)
		}

		if context == "" {
			fmt.Println("プロジェクトコンテキストが設定されていません")
		} else {
			fmt.Printf("現在のプロジェクトコンテキスト: %s\n", context)
		}
		return nil
	},
}

var contextClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear project context",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.SetContext(""); err != nil {
			return fmt.Errorf("failed to clear context: %w", err)
		}
		fmt.Println("プロジェクトコンテキストをクリアしました")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(contextCmd)
	contextCmd.AddCommand(contextSetCmd)
	contextCmd.AddCommand(contextShowCmd)
	contextCmd.AddCommand(contextClearCmd)
}