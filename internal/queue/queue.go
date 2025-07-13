package queue

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Task struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	OutputFiles []string `json:"output_files"`
	UseDeep     bool     `json:"use_deep"`
}

type Queue struct {
	Tasks []Task `json:"tasks"`
}

func getQueuePath() string {
	return filepath.Join(".", ".code_parallel_queue.json")
}

func loadQueue() (*Queue, error) {
	queuePath := getQueuePath()
	
	data, err := os.ReadFile(queuePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Queue{Tasks: []Task{}}, nil
		}
		return nil, err
	}

	var queue Queue
	if err := json.Unmarshal(data, &queue); err != nil {
		return nil, err
	}

	return &queue, nil
}

func saveQueue(queue *Queue) error {
	queuePath := getQueuePath()
	
	data, err := json.MarshalIndent(queue, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(queuePath, data, 0644)
}

func AddTask(task Task) error {
	task.ID = uuid.New().String()
	
	queue, err := loadQueue()
	if err != nil {
		return err
	}

	queue.Tasks = append(queue.Tasks, task)
	return saveQueue(queue)
}

func GetTasks() ([]Task, error) {
	queue, err := loadQueue()
	if err != nil {
		return nil, err
	}

	return queue.Tasks, nil
}

func ClearTasks() error {
	queue := &Queue{Tasks: []Task{}}
	return saveQueue(queue)
}