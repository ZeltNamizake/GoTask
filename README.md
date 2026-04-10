# GoTask ✅️

GoTask is a simple **CLI task manager** built with Go.  
Designed for managing daily tasks quickly and efficiently using JSON-based storage.

---

## ✨ Features

- Add new tasks (`Add Task`)
- Mark tasks as done (`Done Task`)
- Mark tasks as not done (`Undone Task`)
- Delete tasks (`Delete Task`)
- Edit done time on completed tasks (`Edit Done Time`)
- Show tasks by specific date (`ListTasksByDate`)
- Automatically display all tasks on startup
- Store tasks per day (`taskFile/tasks_YYYY-MM-DD.json`)
- Interactive terminal-based interface
- Clean and minimal design, easy to modify

---

## 📁 Project Structure

```plaintext
GoTask/
├── main.go              # CLI entry point
├── go.mod
├── handler/
│   ├── service.go       # Business logic (CRUD, storage)
│   ├── util.go          # Helpers (input, time, clear screen)
│   ├── view.go          # UI rendering (task display)
│   └── task.go          # Task struct
└── taskFile/            # JSON storage (auto-generated)
```

---

## ⚙️ Installation
Make sure **Go** is installed on your system

```bash
git clone https://github.com/username/GoTask.git
cd GoTask
go mod tidy
go run main.go
```

> The taskFile folder will be created automatically when running for the first time, along with an empty JSON file

---

## 🚀 Usage
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
 - ➕️ Add a task:
   ```plaintext
   Title Task: Learning Go
   [INFO] - AddTask successfully
   ```
 - ✅️ Mark as done:
   ```plaintext
   Task Number Done: 1
   [INFO] - Task marked as done ✅️
   ```
 - ↩️ Mark as not done:
   ```plaintext
   Task Number Undone: 1
   [INFO] - Task has been undone ⬜️
   ```
 - ❌️ Delete a task:
   ```plaintext
   Task Number Delete: 1
   [INFO] - Task has been deleted ❌️
   ```

 - 🕑 Time editing on completed tasks:
   ```plaintext
   Task Number: 1
   New Done Time (HH:MM): 08:00
   [INFO] - Done time updated: 08:00 ✅️
   ```

 - 📅 Display task list by date:
   ```plaintext
   Date (YYYY-MM-DD): 2026-04-08

   Tasks [2026-04-08]
   Number of tasks: 3
   1. ✅️ – Learning (07:00) 
   2. ✅️ – Reading (06:00)
   3. ⬜️ – Workout

   [INFO] - Press enter to continue...
   ```

---

## 📝 Notes
- Tasks are stored per day in separate JSON files: taskFile/tasks_YYYY-MM-DD.json
- Input validation is enforced for all fields
- The system uses 0 as a universal escape command
- Output is designed for terminal readability and simplicity

---

## 📜 License
GoTask  is licensed under the MIT License. See the LICENSE file for more information.
