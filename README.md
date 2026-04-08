# **GoTask ✅️**

GoTask is a simple **CLI task manager** built with **Go**.  
Designed for **managing daily tasks** quickly and efficiently, using **JSON-based storage**.  

## **Features**
- Add new tasks (`Add Task`)  
- Mark tasks as done (`Done Task`)  
- Mark tasks as not done (`Undone Task`)  
- Delete tasks (`Delete Task`)  
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
│   ├── service.go      # CRUD tasks & ListTasks
│   └── util.go         # Helpers: input, clear screen, date
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
5. Exit
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
   ```
   Task Number Delete: 1
   [INFO] - Task has been deleted
   ```

## Notes
- Tasks are stored per day, so a new JSON file will be created each day.
- Input must be valid, otherwise an error message will appear.

## License
GoTask  is licensed under the MIT License. See the LICENSE file for more information.
