# GoTask

![Go](https://img.shields.io/github/go-mod/go-version/amahmood561/gotask)
![Build Status](https://img.shields.io/github/actions/workflow/status/amahmood561/gotask/ci.yml?branch=main)

GoTask is a simple command-line task manager built with Go. It allows you to add, list, complete, and delete tasks directly from your terminal, helping you stay organized and productive.

## Features

- **Add Tasks:** Quickly add new tasks with descriptions.
- **List Tasks:** View all your tasks with their statuses.
- **Mark as Complete:** Mark tasks as completed to keep track of your progress.
- **Delete Tasks:** Remove tasks that are no longer needed.
- **Persistent Storage:** Tasks are saved locally in a JSON file, ensuring your data is preserved between sessions.
- **Concurrency Safe:** Handles concurrent operations safely using mutexes.
- **Extensible:** Easily add new features and commands as needed.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)

### Clone the Repository

```bash
git clone https://github.com/amahmood561/GoTask.git
cd gotask
```

### Build the Application

```bash
go build -o gotask
```

### (Optional) Install Globally

To use `gotask` from anywhere in your terminal, you can move the binary to a directory that's in your `PATH`.

```bash
mv gotask /usr/local/bin/
```

## Usage

Once installed, you can use the `gotask` command followed by various subcommands:

### Add a New Task

```bash
gotask add "Buy groceries"
```

### List All Tasks

```bash
gotask list
```

**Sample Output:**

```
[ ] 1: Buy groceries
[X] 2: Read a book
```

### Mark a Task as Completed

```bash
gotask complete 1
```

### Delete a Task

```bash
gotask delete 2
```

### Help

For more information on available commands and usage, use the `--help` flag:

```bash
gotask --help
gotask add --help
```

## Project Structure

```
gotask/
├── cmd/
│   ├── add.go
│   ├── list.go
│   ├── complete.go
│   ├── delete.go
│   └── root.go
├── data/
│   └── tasks.json
├── pkg/
│   └── task/
│       └── task.go
├── tests/
│   └── task_test.go
├── .gitignore
├── README.md
├── go.mod
├── go.sum
└── main.go
```

- **`cmd/`**: Contains CLI command implementations.
- **`data/`**: Stores persistent task data.
- **`pkg/`**: Contains reusable packages and business logic.
- **`tests/`**: Includes unit tests to ensure code reliability.
- **`.gitignore`**: Specifies files and directories to be ignored by Git.
- **`README.md`**: Project documentation.
- **`go.mod` & `go.sum`**: Manage project dependencies.
- **`main.go`**: Entry point of the application.

## Contributing

Contributions are welcome! Whether it's reporting bugs, suggesting new features, or submitting pull requests, your input is valuable.

### How to Contribute

1. **Fork the Repository**

   Click the "Fork" button at the top right of this page to create your own copy of the repository.

2. **Clone the Forked Repository**

   ```bash
   git clone https://github.com/amahmood561/gotask.git
   cd gotask
   ```

3. **Create a New Branch**

   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Make Your Changes**

   Implement your feature or bug fix.

5. **Commit Your Changes**

   ```bash
   git commit -m "Add feature: your feature description"
   ```

6. **Push to Your Fork**

   ```bash
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request**

   Go to the original repository and click the "Compare & pull request" button.

## Contact

For any inquiries or support, please open an issue on [GitHub Issues](https://github.com/amahmood561/gotask/issues) or contact [amahmood561@gmail.com](mailto:amahmood561@gmail.com).

---

*Happy Task Managing!*

---

## Additional Recommendations

### **Add a `.gitignore` File**

Ensure you have a `.gitignore` file to exclude unnecessary files from your Git repository. Here's a sample `.gitignore` tailored for Go projects:

```gitignore
# Binaries
/bin/
/build/
/out/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, build with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE or goland
*.out

# Dependency directories (remove the comment below to include it)
/vendor/

# Go workspace file
go.work

# IDE directories and files
.vscode/
.idea/
*.swp

# Data files
/data/tasks.json

# Logs
*.log
```

### ** todo Set Up Continuous Integration (CI)**

Integrate CI tools like GitHub Actions to automate testing and ensure code quality. example GitHub Actions workflow (`.github/workflows/ci.yml`):

```yaml
name: Go CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:

    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: 1.20

    - name: Install dependencies
      run: go mod tidy

    - name: Run tests
      run: go test ./... -v
```

