package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func ListTasksByDate(date string) error {
	filePath := path.Join(dir, fmt.Sprintf("tasks_%s.json", date))

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			ClearScreen()
			fmt.Printf("[INFO] - No tasks found for this date")
			return nil
		}
		return err
	}

	var tasks []Task
	if len(data) == 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			return err
		}
	}

	ClearScreen()
	fmt.Printf("Tasks [%s]\n", date)
	fmt.Println("Number of tasks:", len(tasks))
	for i, t := range tasks {
		status := "⬜️"
		timeInfo := ""
		if t.Done {
			status = "✅️"
			if t.DoneAt != "" {
				timeInfo = fmt.Sprintf(" (%s)", t.DoneAt)
			}
		}
		fmt.Printf("%d. %s - %s%s\n", i+1, status, t.Title, timeInfo)
	}
	ShowScore()
	return nil

}

func ListTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	fmt.Printf("Tasks [%s]\n", DateNow())
	fmt.Println("Number of tasks:", len(tasks))
	for i, t := range tasks {
		status := "⬜️"
		timeInfo := ""
		if t.Done {
			status = "✅️"
			if t.DoneAt != "" {
				timeInfo = fmt.Sprintf(" (%s)", t.DoneAt)
			}
		}
		fmt.Printf("%d. %s – %s%s\n", i+1, status, t.Title, timeInfo)
	}
	ShowScore()
	return nil
}

func ShowScore() {
	done, total := GetScore()

	percent := 0
	if total > 0 {
		percent = (done * 100) / total
	}
	fmt.Printf("\nProgress: %d%%\n", percent)
	fmt.Printf("Completed Task: %d / %d\n", done, total)
}
