package executor

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"make_parallel/internal/gemini"
	"make_parallel/internal/queue"
)

type TaskResult struct {
	Task   queue.Task
	Error  error
	Status string
}

func ExecuteTasks(tasks []queue.Task, maxWorkers int) error {
	startTime := time.Now()
	
	client, err := gemini.NewClient()
	if err != nil {
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}
	defer client.Close()

	taskChan := make(chan queue.Task, len(tasks))
	resultChan := make(chan TaskResult, len(tasks))
	
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	var wg sync.WaitGroup
	
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(client, taskChan, resultChan, &wg, i+1)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	createdFiles := []string{}
	failedTasks := 0
	
	for result := range resultChan {
		if result.Error != nil {
			fmt.Printf("❌ Task failed: %s - %v\n", result.Task.Description, result.Error)
			failedTasks++
		} else {
			fmt.Printf("✅ Task completed: %s\n", result.Task.Description)
			for _, file := range result.Task.OutputFiles {
				createdFiles = append(createdFiles, file)
			}
		}
	}

	duration := time.Since(startTime)
	
	fmt.Printf("\n実行完了:\n")
	if len(createdFiles) > 0 {
		fmt.Printf("生成されたファイル:\n")
		for _, file := range createdFiles {
			if fileExists(file) {
				fmt.Printf("- %s (新規作成)\n", file)
			}
		}
	}
	
	if failedTasks > 0 {
		fmt.Printf("失敗したタスク: %d件\n", failedTasks)
	}
	
	fmt.Printf("実行時間: %.1f秒\n", duration.Seconds())
	fmt.Printf("並列数: %d\n", maxWorkers)

	return nil
}

func worker(client *gemini.Client, taskChan <-chan queue.Task, resultChan chan<- TaskResult, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	
	for task := range taskChan {
		fmt.Printf("🔄 [Worker %d] Processing: %s\n", workerID, task.Description)
		
		result := processTask(client, task, workerID)
		resultChan <- result
	}
}

func processTask(client *gemini.Client, task queue.Task, workerID int) TaskResult {
	ctx := context.Background()
	
	for attempt := 1; attempt <= 2; attempt++ {
		if attempt > 1 {
			fmt.Printf("🔄 [Worker %d] Retrying task: %s (attempt %d)\n", workerID, task.Description, attempt)
		}
		
		content, err := client.GenerateCode(ctx, task.Description, task.UseDeep)
		if err != nil {
			if attempt == 2 {
				return TaskResult{
					Task:   task,
					Error:  fmt.Errorf("failed after retry: %w", err),
					Status: "failed",
				}
			}
			time.Sleep(2 * time.Second)
			continue
		}

		for _, outputFile := range task.OutputFiles {
			if err := saveToFile(outputFile, content); err != nil {
				if attempt == 2 {
					return TaskResult{
						Task:   task,
						Error:  fmt.Errorf("failed to save file %s: %w", outputFile, err),
						Status: "failed",
					}
				}
				time.Sleep(1 * time.Second)
				break
			}
		}

		return TaskResult{
			Task:   task,
			Error:  nil,
			Status: "completed",
		}
	}

	return TaskResult{
		Task:   task,
		Error:  fmt.Errorf("unexpected error"),
		Status: "failed",
	}
}

func saveToFile(filePath, content string) error {
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return os.WriteFile(filePath, []byte(content), 0644)
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}