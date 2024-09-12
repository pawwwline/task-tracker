# Task Tracker
***
Simple command line task management app. Implemented in Goland and uses JSON file as a simple database.

## Features
- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

## Installation

- Ensure you have Golang installed

1. Clone repository:
```bash
git clone https://github.com/pawwwline/task-tracker
```
2. Navigate to project
```bash
cd task-tracker
```

3. Build the application
```bash
go build task-tracker
```

## Usage
***
### Add task
it will return id of added task
```bash
./task-tracker add "Buy groceries"
```
### Update task
update task by id number (you can check it using list command)
```bash
./task-tracker update 1 "Clean room"
```
### Delete task
delete task by its id number
```bash
./task-tracker delete 1
```
***
### Mark tasks
#### Mark done
```bash
./task-tracker mark-done 1
```
#### Mark in-progress
```bash
./task-tracker mark-in-progress 1
```

#### Mark to-do
```bash
./task-tracker mark-todo 1
```
***
### List tasks
#### List all tasks
```bash
./task-tracker list
```
#### List to-do tasks
```bash
./task-tracker list todo
```
#### List in progress tasks
```bash
./task-tracker list in-progress
```
#### List done tasks
```bash
./task-tracker list done
```
***
project challenge
[task-tracker](https://roadmap.sh/projects/task-tracker) by
[roadmap.sh](https://roadmap.sh/)
