# GoTask ✅️

GoTask is a lightweight CLI tool for tracking real task execution.

Instead of managing plans, GoTask focuses on logging what you actually complete — along with timestamps — using simple JSON-based storage.

---

## 🧠 Philosophy

GoTask is not a traditional todo list.

It is designed to:
- Track what you actually complete, not what you plan
- Provide clear execution logs with timestamps
- Keep the system minimal and distraction-free

No backlog. No over-planning.
Only real work that has been done.

---

## ✨ Features

- Add new tasks (`Add Task`)
- Mark tasks as done (`Done Task`)
- Mark tasks as not done (`Undone Task`)
- Delete tasks (`Delete Task`)
- Edit done time on completed tasks (`Edit Done Time`)
- Edit start time on task started (`Edit Start Time`)
- Edit the task title (`Edit Title Task`)
- Show tasks by specific date (`ListTasksByDate`)
- Show available task dates (`Show Available Dates`)  
- Display task completion score (Done / Total + Progress %)  
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
│   ├── util.go          # Helpers (input, score/progress, clear screen)
│   ├── view.go          # UI rendering (task display)
│   ├── time_engine.go   # setup timezone (first-run setup)
│   └── task.go          # Task struct
├── config/              # Configuration (auto-generated)
│   └── timezone.json    # User timezone setting (first-run setup)
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

> The `taskFile` folder is automatically created on first run and is used as the primary storage for all task data, separated by date in JSON format.

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
6. Edit Start Time
7. Edit Title Task
8. Show Tasks by Date
9. Show Available Dates
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
 - 🕑 Time editing on task started:
   ```plaintext
   Task Number: 1
   New Start Time (HH:MM): 07:00
   [INFO] - Start time updated: 07:00 ✅️
   ```

 - ✏️ Change the task title:
   ```plaintext
   Task Number: 1
   New Title: Learning
   [INFO] - Edit task title successfully ✅️
   ```

 - 📅 Display task list by date:
   ```plaintext
   Date (YYYY-MM-DD): 2026-04-08

   📅 Tasks [2026-04-08]
   ────────────────────────────────────────────
   No  Time            Status Title
   ────────────────────────────────────────────
   1   07:00 - 08:00   ✅️     Learning
   2   08:00 - 08:30   ✅️     Reading
   3   09:00           ⬜️     Workout
   ────────────────────────────────────────────

   Progress: 66%
   Completed Task: 2 / 3

   [INFO] - Press enter to continue...
   ```

 - 📅 Show available dates:
   ```plaintext
   📅 Available Dates:
   ───────────────────────
   No   Date
   ───────────────────────
   1    2026-04-09
   2    2026-04-10
   3    2026-04-12
   4    2026-04-14
   ───────────────────────
   [INFO] - Press enter to continue...
   ```

---

## 📝 Notes
- Tasks are stored per day in separate JSON files: taskFile/tasks_YYYY-MM-DD.json
- Input validation is enforced for all fields
- The system uses 0 as a universal escape command
- Output is designed for terminal readability and simplicity
- Score is calculated based on completed tasks
- Available dates are automatically detected from existing task files
- Time handling is consistent across platforms using timezone configuration (`config/timezone.json`).
- The title of the task can be changed if you want to update it

---

## 📜 License
GoTask  is licensed under the MIT License. See the LICENSE file for more information.
