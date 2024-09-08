# Go ToDo CLI Application

A simple command-line ToDo application written in Go. This app allows you to manage your ToDo list with features such as adding, deleting, toggling, editing, and listing tasks.

![Screenshot 2024-09-08 125736](https://github.com/user-attachments/assets/0fdcec39-7410-43f0-8168-8675e45bb43c)


## Features

- Add new tasks
- Delete existing tasks
- Toggle the completion status of tasks
- Edit tasks
- List all tasks with detailed information

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)
- [License](#license)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/raghulchandramouli/CLI-ToDo-Application
    cd CLI-ToDo-Application
    ```

2. Build the project:
    ```bash
    go build -o CLI-ToDo-Application
    ```

3. Run the application:
    ```bash
    ./CLI-ToDo-Application
    ```

## Usage

This CLI app uses flags for various operations. Below are the supported commands:

### Commands:

- `-add <title>`: Adds a new ToDo item with the specified title.
- `-Del <index>`: Deletes the ToDo at the given index.
- `-Toggle <index>`: Toggles the completion status of the ToDo at the given index.
- `-Edit <id:new_title>`: Edits the title of the ToDo at the specified index.
- `-List`: Lists all ToDos.

### Examples:

1. **Add a new ToDo**:
    ```bash
    ./todo-app -add "Buy groceries"
    ```

2. **List all ToDos**:
    ```bash
    ./todo-app -List
    ```

3. **Delete a ToDo**:
    ```bash
    ./todo-app -Del 2
    ```

4. **Toggle the completion status**:
    ```bash
    ./todo-app -Toggle 1
    ```

5. **Edit a ToDo**:
    ```bash
    ./todo-app -Edit 2:"Pick up laundry"
    ```

### Flag Definitions:

- `-add <title>`: Add a new task with the specified title.
- `-Del <index>`: Delete the task at the specified index.
- `-Toggle <index>`: Toggle the completion status of the task at the specified index.
- `-Edit <index>:<new_title>`: Edit the task's title at the given index.
- `-List`: List all current tasks with their statuses.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
