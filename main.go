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
5. Show Tasks by Date
6. Exit
Select: `

	for {
		handler.ClearScreen()
		if err := handler.ListTasks(); err != nil {
			fmt.Println("[ERR] - ListTasks:", err)
		}

		selected := handler.AskRequired(menu)
		switch selected {

		case "1":
			titleTask := handler.AskRequired("Title Task: ")
			if err := handler.AddTask(titleTask); err != nil {
				fmt.Println("[ERR] - AddTask:", err)
			} else {
				fmt.Println("[INFO] - AddTask successfully")
			}

		case "2":
			index := handler.AskIntRequired("Task Number Done: ")

			if err := handler.DoneTask(index); err != nil {
				fmt.Println("[ERR] - DoneTask:", err)
			} else {
				fmt.Println("[INFO] - Task marked as done ✅️")	
			}

		case "3":
			index := handler.AskIntRequired("Task Number Undone: ")

			if err := handler.UndoneTask(index); err != nil {
				fmt.Println("[ERR] - UndoneTask:", err)
			} else {
				fmt.Println("[INFO] - Task has been undone")
			}

		case "4":
			index := handler.AskIntRequired("Task Number Delete: ")

			if err := handler.DeleteTask(index); err != nil {
				fmt.Println("[ERR] - DeleteTask:", err)
			} else {
				fmt.Println("[INFO] - Task has been deleted")
			}

		case "5":
			date := handler.AskRequired("Date (YYYY-MM-DD): ")

			if err := handler.ListTasksByDate(date); err != nil {
				fmt.Println("[ERR] - ListTasksByDate:", err)
			} else {
				handler.Ask("\n[INFO] - Press enter to continue...")
			}

		case "6":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("[ERR] - Invalid choice")
		}
	}

}
