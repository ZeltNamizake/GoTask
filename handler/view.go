package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ListTasksByDate(date string) error {
	filePath := filepath.Join(dir, fmt.Sprintf("tasks_%s.json", date))
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
	fmt.Println("────────────────────────────────────────────")
	fmt.Printf("%-3s %-15s %-6s %s\n", "No", "Time", "Status", "Title")
	fmt.Println("────────────────────────────────────────────")

	for i, t := range tasks {
		status := "⬜️"
		timeInfo := ""

		if t.CreateAt != "" && t.DoneAt != "" {
			timeInfo = fmt.Sprintf("%s - %s", t.CreateAt, t.DoneAt)
		} else if t.DoneAt != "" {
			timeInfo = t.DoneAt
		} else if t.CreateAt != "" {
			timeInfo = t.CreateAt
		}

		if t.Done {
			status = "✅️"
		}

		title := Truncate(t.Title, 40)
		fmt.Printf("%-3d %-15s %-6s %s\n",
			i+1, timeInfo, status, title)
	}

	fmt.Println("────────────────────────────────────────────")
	ShowScoreByDate(date)
	return nil
}

func ListTasks() error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}
	fmt.Printf("📅 Tasks [%s]\n", DateNow())
	fmt.Println("────────────────────────────────────────────")
	fmt.Printf("%-3s %-15s %-6s %s\n", "No", "Time", "Status", "Title")
	fmt.Println("────────────────────────────────────────────")

	for i, t := range tasks {
		status := "⬜️"
		timeInfo := ""

		if t.CreateAt != "" && t.DoneAt != "" {
			timeInfo = fmt.Sprintf("%s - %s", t.CreateAt, t.DoneAt)
		} else if t.DoneAt != "" {
			timeInfo = t.DoneAt
		} else if t.CreateAt != "" {
			timeInfo = t.CreateAt
		}

		if t.Done {
			status = "✅️"
		}

		title := Truncate(t.Title, 40)
		fmt.Printf("%-3d %-15s %-6s %s\n",
			i+1, timeInfo, title, status)
	}

	fmt.Println("────────────────────────────────────────────")
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
	fmt.Println("───────────────────────")
	fmt.Printf("%-4s %-5s", "No", "Date\n")
	fmt.Println("───────────────────────")

	for i, date := range dates {
		fmt.Printf("%-4d %-5s\n", i+1, date)
	}

	fmt.Println("───────────────────────")
	return nil

}

func Truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-1] + "..."
	}
	return s
}
