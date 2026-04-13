package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ListTasksByDate(date string) error {
	filePath := filepath.Join(dir, fmt.Sprintf("tasks_%s.json", date))

	fmt.Println(filePath)

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
		ClearScreen()
		fmt.Println("[INFO] - Empty file.")
		return nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	ClearScreen()
	fmt.Printf("📅 Tasks [%s]\n", date)
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
		index := fmt.Sprintf("%d.", i+1)
		fmt.Printf("%-3s %s - %s%s\n", index, status, t.Title, timeInfo)
	}
	ShowScoreByDate(date)
	return nil

}

func ListTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	fmt.Printf("📅 Tasks [%s]\n", DateNow())
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
		index := fmt.Sprintf("%d.", i+1)
		fmt.Printf("%-3s %s – %s%s\n", index, status, t.Title, timeInfo)
	}
	ShowScoreNow()
	return nil
}

func ShowScoreNow() {
	done, total := GetScoreNow()

	percent := 0
	if total > 0 {
		percent = (done * 100) / total
	}
	fmt.Printf("\nProgress: %d%%\n", percent)
	fmt.Printf("Completed Task: %d / %d\n", done, total)
}

func ShowScoreByDate(date string) {
	done, total := GetScoreByDate(date)

	percent := 0
	if total > 0 {
		percent = (done * 100) / total
	}
	fmt.Printf("\nProgress: %d%%\n", percent)
	fmt.Printf("Completed Task: %d / %d\n", done, total)
}

func ShowAvailableDates() error {
	dates, err := GetAvailableDates()
	if err != nil {
		return err
	}

	if len(dates) == 0 {
		fmt.Println("[INFO] - No task dates found")
		return nil
	}

	fmt.Println("📅 Available Dates:")
	for i, date := range dates {
		fmt.Printf("%d. %s\n", i+1, date)
	}
	return nil

}
