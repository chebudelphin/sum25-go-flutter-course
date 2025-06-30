package taskmanager

import (
	"errors"
	"time"
)

// Predefined errors
var (
	ErrTaskNotFound = errors.New("task not found")
<<<<<<< HEAD
	ErrEmptyTitle   = errors.New("task title cannot be empty")
	ErrInvalidID    = errors.New("invalid task ID")
=======
	ErrEmptyTitle   = errors.New("title cannot be empty")
>>>>>>> e6d76f7 (update lab1 and workflow of submission)
)

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}
type TaskManager struct {
	tasks  map[int]Task
	nextID int
}

func NewTaskManager() *TaskManager {
<<<<<<< HEAD
	return &TaskManager{
		tasks:  make(map[int]*Task),
		nextID: 1,
	}
}
func (tm *TaskManager) AddTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}
	task := &Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}
	tm.tasks[tm.nextID] = task
	tm.nextID++
	return task, nil
}
func (tm *TaskManager) UpdateTask(id int, title, description string, done bool) error {
	if id <= 0 {
		return ErrInvalidID
	}
	task, exists := tm.tasks[id]
	if !exists {
		return ErrTaskNotFound
	}
	if title == "" {
		return ErrEmptyTitle
	}
	task.Title = title
	task.Description = description
	task.Done = done
	return nil
}
func (tm *TaskManager) DeleteTask(id int) error {
	if id <= 0 {
		return ErrInvalidID
	}
	if _, exists := tm.tasks[id]; !exists {
		return ErrTaskNotFound
	}
	delete(tm.tasks, id)
	return nil
}
func (tm *TaskManager) GetTask(id int) (*Task, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}
	task, exists := tm.tasks[id]
	if !exists {
		return nil, ErrTaskNotFound
	}
	return task, nil
}
func (tm *TaskManager) ListTasks(filterDone *bool) []*Task {
	result := []*Task{}
	for _, task := range tm.tasks {
		if filterDone == nil {
			result = append(result, task)
		} else {
			if task.Done == *filterDone {
				result = append(result, task)
			}
		}
	}
	return result
=======
	// TODO: Implement this function
	return nil
}

// AddTask adds a new task to the manager, returns an error if the title is empty, and increments the nextID
func (tm *TaskManager) AddTask(title, description string) (Task, error) {
	// TODO: Implement this function
	return Task{}, nil
}

// UpdateTask updates an existing task, returns an error if the title is empty or the task is not found
func (tm *TaskManager) UpdateTask(id int, title, description string, done bool) error {
	// TODO: Implement this function
	return nil
}

// DeleteTask removes a task from the manager, returns an error if the task is not found
func (tm *TaskManager) DeleteTask(id int) error {
	// TODO: Implement this function
	return nil
}

// GetTask retrieves a task by ID, returns an error if the task is not found
func (tm *TaskManager) GetTask(id int) (Task, error) {
	// TODO: Implement this function
	return Task{}, nil
}

// ListTasks returns all tasks, optionally filtered by done status, returns an empty slice if no tasks are found
func (tm *TaskManager) ListTasks(filterDone *bool) []Task {
	// TODO: Implement this function
	return nil
>>>>>>> e6d76f7 (update lab1 and workflow of submission)
}
