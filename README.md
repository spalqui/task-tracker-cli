# TaskTracker CLI
TaskTracker CLI is a simple but yet effective way to manage tasks in the current working directory. It will persist the tasks
in a JSON file named `tasks.json`.

## Usage

Add a new task.
```bash
./task-cli add "buy milk"
```

Update a task.
```bash
./task-cli update 1 "buy milk x2"
```

List all tasks.
```bash
./task-cli list
```

List all todo tasks.
```bash
./task-cli list todo
```

List all in-progress tasks.
```bash
./task-cli list in-progress
```

List all done tasks.
```bash
./task-cli list in-progress
```

Delete a task.
```bash
./task-cli delete 1
```

### TODOs
- Implement the command pattern for better extensibility when adding new commands in the future.
