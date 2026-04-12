package main

import (
	"fmt"

	"GoTask/handler"
)

func main() {
	if err := handler.InitStorage(); err != nil {
		fmt.Println("InitStorage:", err)
		return
	}

	menu := `
1. Add Task
2. Done Task
3. Undone Task
4. Delete Task
5. Edit Done Time
6. Show Tasks by Date
7. Show Available Dates
0. Back/Exit
Select: `

	for {
		handler.ClearScreen()
		if err := handler.ListTasks(); err != nil {
			fmt.Println("[ERR] - ListTasks:", err)
		}

		selected, ok := handler.AskRequired(menu)
		if !ok {
			fmt.Println("[INFO] - Exit")
			return
		}

		switch selected {

		case "1":
			titleTask, ok := handler.AskRequired("Title Task: ")
			if !ok {
				fmt.Println("[INFO] - Cancelled")
				continue
			}

			if err := handler.AddTask(titleTask); err != nil {
				fmt.Println("[ERR] - AddTask:", err)
			} else {
				fmt.Println("[INFO] - AddTask successfully")
			}

		case "2":
			index, ok := handler.AskIntRequired("Task Number Done: ")
			if !ok {
				fmt.Println("[INFO] - Cancelled")
				continue
			}

			if err := handler.DoneTask(index); err != nil {
				fmt.Println("[ERR] - DoneTask:", err)
			} else {
				fmt.Println("[INFO] - Task marked as done ✅️")	
			}

		case "3":
			index, ok := handler.AskIntRequired("Task Number Undone: ")
			if !ok {
				fmt.Println("[INFO] - Cancelled")
				continue
			}

			if err := handler.UndoneTask(index); err != nil {
				fmt.Println("[ERR] - UndoneTask:", err)
			} else {
				fmt.Println("[INFO] - Task has been undone ⬜️")
			}

		case "4":
			index, ok := handler.AskIntRequired("Task Number Delete: ")
			if !ok {
				fmt.Println("[INFO] - Cancelled")
				continue
			}

			if err := handler.DeleteTask(index); err != nil {
				fmt.Println("[ERR] - DeleteTask:", err)
			} else {
				fmt.Println("[INFO] - Task has been deleted ❌️")
			}

		case "5":
			index, ok := handler.AskIntRequired("Task Number: ")
			if !ok {
				fmt.Println("[INFO] - Cancelled")
				continue
			}

			input, ok := handler.AskRequired("New Done Time (HH:MM): ")
			if !ok {
				fmt.Println("[INFO] - Cancelled")
				continue
			}

			time, err := handler.ValidateTime(input)
			if err != nil {
				fmt.Println("[ERR] -", err)
				continue
			}

			if err := handler.EditDoneTime(index, time); err != nil {
				fmt.Println("[ERR] - EditDoneTime:", err)
			} else {
				fmt.Println("[INFO] - Done time updated:", time, "✅️")
			}

		case "6":
			date, ok := handler.AskRequired("Date (YYYY-MM-DD): ")
			if !ok {
				fmt.Println("[INFO] - Cancelled")
				continue
			}

			if err := handler.ListTasksByDate(date); err != nil {
				fmt.Println("[ERR] - ListTasksByDate:", err)
			} else {
				handler.Ask("\n[INFO] - Press enter to continue...")
			}

		case "7":
			handler.ClearScreen()
			if err := handler.ShowAvailableDates(); err != nil {
				fmt.Println("[ERR] - ShowAvailableDates:", err)
			} else {
				handler.Ask("\n[INFO] - Press enter to continue...")
			}

		default:
			fmt.Println("[ERR] - Invalid choice")
		}
	}

}
