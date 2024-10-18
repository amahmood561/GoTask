package task

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "os"
    "sync"
)

const dataFile = "data/tasks.json"

type Task struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}

var tasks []Task
var mu sync.Mutex

func init() {
    loadTasks()
}

func loadTasks() {
    mu.Lock()
    defer mu.Unlock()

    if _, err := os.Stat(dataFile); os.IsNotExist(err) {
        tasks = []Task{}
        saveTasks()
        return
    }

    data, err := ioutil.ReadFile(dataFile)
    if err != nil {
        panic(err)
    }

    json.Unmarshal(data, &tasks)
}

func saveTasks() {
    mu.Lock()
    defer mu.Unlock()

    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        panic(err)
    }

    os.MkdirAll("data", os.ModePerm)
    ioutil.WriteFile(dataFile, data, 0644)
}

func GetAllTasks() []Task {
    mu.Lock()
    defer mu.Unlock()
    return tasks
}

func AddTask(t Task) {
    mu.Lock()
    defer mu.Unlock()
    tasks = append(tasks, t)
    saveTasks()
}

func GetNextID() int {
    mu.Lock()
    defer mu.Unlock()
    max := 0
    for _, t := range tasks {
        if t.ID > max {
            max = t.ID
        }
    }
    return max + 1
}

func CompleteTask(id int) error {
    mu.Lock()
    defer mu.Unlock()
    for i, t := range tasks {
        if t.ID == id {
            tasks[i].Completed = true
            saveTasks()
            return nil
        }
    }
    return errors.New("task not found")
}

func DeleteTask(id int) error {
    mu.Lock()
    defer mu.Unlock()
    for i, t := range tasks {
        if t.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            saveTasks()
            return nil
        }
    }
    return errors.New("task not found")
}
