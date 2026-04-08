package handler

import (
	"encoding/json"
	"fmt"
	"os"
)

func InitStorage() error {
	path := GetPath()

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		empty := []Task{}
		data, err := json.MarshalIndent(empty, "", "  ")
		if err != nil {
			return err
		}

		if err := os.WriteFile(path, data, 0644); err != nil {
			return err
		}
	}
	return nil
}

func LoadTasks() ([]Task, error) {
	path := GetPath()

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tasks []Task

	if len(data) == 0 {
		return []Task{}, nil
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return []Task{}, nil
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	path := GetPath()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}

func AddTask(title string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, Task{
		Title: title,
		Done:  false,
	})

	return SaveTasks(tasks)
}

func DoneTask(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("Invalid index")
	}

	if tasks[index-1].Done {
		return fmt.Errorf("Task is already done")
	}

	tasks[index-1].Done = true
	return SaveTasks(tasks)
}

func UndoneTask(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("Invalid index")
	}

	if !tasks[index-1].Done {
		return fmt.Errorf("Task is already not done")
	}

	tasks[index-1].Done = false
	return SaveTasks(tasks)
}

func DeleteTask(index int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	if index < 1 || index > len(tasks) {
		return fmt.Errorf("Invalid index")
	}

	tasks = append(tasks[:index-1], tasks[index:]...)

	return SaveTasks(tasks)
}

func ListTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	fmt.Printf("Tasks [%s]\n", DateNow())
	fmt.Println("Number of tasks:", len(tasks))
	for i, t := range tasks {
		status := "❌️"
		if t.Done {
			status = "✅️"
		}
		fmt.Printf("%d. %s – %s\n", i+1, status, t.Title)
	}
	return nil
}
