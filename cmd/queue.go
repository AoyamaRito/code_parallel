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
	Use:   "queue [task-json]",
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

		fmt.Printf("Task added: %s\n", description)
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
			fmt.Println("No tasks in queue")
			return nil
		}

		fmt.Printf("Executing %d tasks with %d workers...\n", len(tasks), parallel)
		
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
			fmt.Println("No tasks in queue")
			return nil
		}

		fmt.Printf("Queue contains %d tasks:\n", len(tasks))
		for i, task := range tasks {
			model := "Gemini 2.0 Flash"
			if task.UseDeep {
				model = "Gemini 2.0 Pro"
			}
			fmt.Printf("%d. %s (%s)\n", i+1, task.Description, model)
			for _, file := range task.OutputFiles {
				fmt.Printf("   -> %s\n", file)
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
		fmt.Println("Queue cleared")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(queueAddCmd)
	queueCmd.AddCommand(queueRunCmd)
	queueCmd.AddCommand(queueListCmd)
	queueCmd.AddCommand(queueClearCmd)
	
	queueRunCmd.Flags().IntP("parallel", "p", 10, "Number of parallel workers")
}