# **GoTask ✅️**

GoTask is a simple **CLI task manager** built with **Go**.  
Designed for **managing daily tasks** quickly and efficiently, using **JSON-based storage**.  

## **Features**
- Add new tasks (`Add Task`)  
- Mark tasks as done (`Done Task`)  
- Mark tasks as not done (`Undone Task`)  
- Delete tasks (`Delete Task`)  
- Display tasks for a specific date (`ListTasksByDate`)  
- Edit time on completed tasks (`EditDoneTime`)  
- In input fields → cancel current action  
- In main menu → exit application  
- Automatically display all tasks  
- Stores tasks by date (`taskFile/tasks_<date>.json`)  
- Interactive terminal interface  
- Clean and minimal, easy to modify  

## **Project Structure**
```plaintext
GoTask/
├── main.go             # CLI entry point
├── go.mod
├── handler/
│   ├── service.go      # CRUD LoadTask, InitStorage, etc
│   └── util.go         # Helpers: input, clear screen, date
|   └── view.go         # View: List task
└── taskFile/           # Folder for storing JSON task files
```

## Installation
Make sure **Go** is installed on your system

```bash
git clone https://github.com/username/GoTask.git
cd GoTask
go mod tidy
go run main.go
```

> The taskFile folder will be created automatically when running for the first time, along with an empty JSON file

## Usage
1. Run the program:
```bash
go run main.go
```

2. Select a menu option:
```plaintext
1. Add Task
2. Done Task
3. Undone Task
4. Delete Task
5. Edit Done Time
6. Show Tasks by Date
0. Back/Exit
Select:
```

3. Follow the instructions according to the menu. For example:
 - Add a task:
   ```plaintext
   Title Task: Learning Go
   [INFO] - AddTask successfully
   ```
 - Mark as done:
   ```plaintext
   Task Number Done: 1
   [INFO] - Task marked as done ✅️
   ```
 - Mark as not done:
   ```plaintext
   Task Number Undone: 1
   [INFO] - Task has been undone
   ```
 - Delete a task:
   ```plaintext
   Task Number Delete: 1
   [INFO] - Task has been deleted
   ```

 - Time editing on completed tasks:
   ```plaintext
   Task Number Markdone: 1
   New Done Time (HH:MM): 08:00
   [INFO] - Done time update: 08:00
   ```

 - Display task list by date:
   ```plaintext
   Date (YYYY-MM-DD): 2026-04-08

   Tasks [2026-04-08]
   Number of tasks: 3
   1. ✅️ – Learning (08.00) 
   2. ✅️ – Reading (06.00)
   3. ⬜️ – Workout

   [INFO] - Press enter to continue...
   ```

## Notes
- Tasks are stored per day, so a new JSON file will be created each day.
- Input must be valid, otherwise an error message will appear.

## License
GoTask  is licensed under the MIT License. See the LICENSE file for more information.
