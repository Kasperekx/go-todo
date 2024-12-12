package storage

import (
	"fmt"
	"time"

	"github.com/kasperekx/todo/internal/models"
)

type TodoList struct {
	todos  map[int]models.Todo
	nextID int
}

func NewToDoList() *TodoList {
	return &TodoList{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

func (tl *TodoList) AddTodo(name string) models.Todo {
	todo := models.Todo{
		ID:        tl.nextID,
		Name:      name,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tl.todos[tl.nextID] = todo
	tl.nextID++
	return todo

}

func (tl *TodoList) ShowAllTodos() []models.Todo {
	todos := make([]models.Todo, 0, len(tl.todos))
	for _, todo := range tl.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (tl *TodoList) RemoveTodo(id int) bool {
	if _, exists := tl.todos[id]; !exists {
		fmt.Printf("There is no todo with ID %d\n", id)
		return false
	}

	delete(tl.todos, id)
	return true

}

func (tl *TodoList) MarkAsCompleted(id int) bool {
	todo, exists := tl.todos[id]
	if !exists {
		return false
	}
	todo.Completed = true
	tl.todos[id] = todo
	return true
}
