package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"make_parallel/internal/queue"
	"make_parallel/internal/executor"
)

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Manage task queue",
}

var queueAddCmd = &cobra.Command{
	Use:   "[task-json]",
	Short: "Add task to queue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var taskArgs []string
		if err := json.Unmarshal([]byte(args[0]), &taskArgs); err != nil {
			return fmt.Errorf("invalid JSON format: %w", err)
		}

		if len(taskArgs) < 2 {
			return fmt.Errorf("task must have at least description and one output file")
		}

		description := taskArgs[0]
		outputs := taskArgs[1:]
		
		useDeep := false
		if len(outputs) > 0 && outputs[len(outputs)-1] == "deep" {
			useDeep = true
			outputs = outputs[:len(outputs)-1]
		}

		task := queue.Task{
			Description: description,
			OutputFiles: outputs,
			UseDeep:     useDeep,
		}

		if err := queue.AddTask(task); err != nil {
			return fmt.Errorf("failed to add task: %w", err)
		}

		tasks, _ := queue.GetTasks()
		fmt.Printf("キューに追加しました: [%d] %s\n", len(tasks), description)
		return nil
	},
}

var queueRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute queued tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		parallel, _ := cmd.Flags().GetInt("parallel")
		
		tasks, err := queue.GetTasks()
		if err != nil {
			return fmt.Errorf("failed to get tasks: %w", err)
		}

		if len(tasks) == 0 {
			fmt.Println("キュー内にタスクがありません")
			return nil
		}

		fmt.Printf("コード生成を開始します... (タスク数: %d, 並列数: %d)\n", len(tasks), parallel)
		
		if err := executor.ExecuteTasks(tasks, parallel); err != nil {
			return fmt.Errorf("execution failed: %w", err)
		}

		if err := queue.ClearTasks(); err != nil {
			return fmt.Errorf("failed to clear tasks: %w", err)
		}

		return nil
	},
}

var queueListCmd = &cobra.Command{
	Use:   "list",
	Short: "List queued tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := queue.GetTasks()
		if err != nil {
			return fmt.Errorf("failed to get tasks: %w", err)
		}

		if len(tasks) == 0 {
			fmt.Println("キュー内のタスク数: 0")
			return nil
		}

		fmt.Printf("キュー内のタスク数: %d\n\n", len(tasks))
		for i, task := range tasks {
			model := "Gemini 2.0 Flash Lite"
			if task.UseDeep {
				model = "Gemini 2.0 Flash (deep)"
			}
			fmt.Printf("[%d] %s\n", i+1, task.Description)
			fmt.Printf("    出力ファイル: %s\n", task.OutputFiles[0])
			if len(task.OutputFiles) > 1 {
				for _, file := range task.OutputFiles[1:] {
					fmt.Printf("                 %s\n", file)
				}
			}
			fmt.Printf("    モデル: %s\n", model)
			if i < len(tasks)-1 {
				fmt.Println()
			}
		}
		return nil
	},
}

var queueClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all queued tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := queue.ClearTasks(); err != nil {
			return fmt.Errorf("failed to clear tasks: %w", err)
		}
		fmt.Println("キューをクリアしました")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(queueCmd)
	queueCmd.RunE = queueAddCmd.RunE
	queueCmd.Args = queueAddCmd.Args
	queueCmd.AddCommand(queueRunCmd)
	queueCmd.AddCommand(queueListCmd)
	queueCmd.AddCommand(queueClearCmd)
	
	queueRunCmd.Flags().IntP("parallel", "p", 10, "Number of parallel workers")
}