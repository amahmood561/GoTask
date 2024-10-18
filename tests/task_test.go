package tests

import (
    "os"
    "testing"

    "github.com/yourusername/gotask/pkg/task"
)

func TestAddTask(t *testing.T) {
    // Setup: Ensure a clean state
    os.Remove("data/tasks.json")
    task.LoadTasks()

    newTask := task.Task{
        Description: "Test Task",
        Completed:   false,
    }

    task.AddTask(newTask)
    tasks := task.GetAllTasks()

    if len(tasks) != 1 {
        t.Errorf("Expected 1 task, got %d", len(tasks))
    }

    if tasks[0].Description != "Test Task" {
        t.Errorf("Expected task description 'Test Task', got '%s'", tasks[0].Description)
    }
}

func TestCompleteTask(t *testing.T) {
    // Setup
    os.Remove("data/tasks.json")
    task.LoadTasks()

    newTask := task.Task{
        Description: "Complete Task",
        Completed:   false,
    }

    task.AddTask(newTask)
    err := task.CompleteTask(1)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    tasks := task.GetAllTasks()
    if !tasks[0].Completed {
        t.Errorf("Expected task to be completed")
    }
}

func TestDeleteTask(t *testing.T) {
    // Setup
    os.Remove("data/tasks.json")
    task.LoadTasks()

    newTask := task.Task{
        Description: "Delete Task",
        Completed:   false,
    }

    task.AddTask(newTask)
    err := task.DeleteTask(1)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    tasks := task.GetAllTasks()
    if len(tasks) != 0 {
        t.Errorf("Expected 0 tasks, got %d", len(tasks))
    }
}
