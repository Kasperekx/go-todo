package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kasperekx/todo/helper"
	"github.com/kasperekx/todo/internal/storage"
)

type Handler struct {
	todoList *storage.TodoList
	options  []string
}

func NewHandler(todoList *storage.TodoList) *Handler {
	return &Handler{
		todoList: todoList,
		options: []string{
			"1. Add todo",
			"2. Show all todos",
			"3. Remove todo",
			"4. Mark as completed",
			"5. Exit",
		},
	}
}

func (h *Handler) Run() {
	var choice int

	for {
		fmt.Println("========= TO DO CLI =========")
		helper.ShowOptions(h.options)
		fmt.Println("=============================")
		fmt.Print("What you want to do: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			h.handleAddTodo()
		case 2:
			h.handleShowTodos()
		case 3:
			h.handleRemoveTodo()
		case 4:
			h.handlerComplete()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Incorrect option")
		}
	}
}

func (h *Handler) handleAddTodo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name of your todo: ")

	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("Todo name cannot be empty string!")
		return
	}

	todo := h.todoList.AddTodo(name)
	fmt.Printf("Rodo %s added successfully with ID: %d\n", todo.Name, todo.ID)
}

func (h *Handler) handleShowTodos() {
	todos := h.todoList.ShowAllTodos()
	fmt.Println("List of todos")
	if len(todos) == 0 {
		fmt.Println("There is no available todo!")
	}
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func (h *Handler) handleRemoveTodo() {
	if len(h.todoList.ShowAllTodos()) == 0 {
		fmt.Println("No todos to remove.")
		return
	}

	fmt.Println("Current todos: ")
	h.handleShowTodos()

	var id int

	for {
		fmt.Println("Please give id of todo you want to remove: ")
		_, err := fmt.Scan(&id)

		if err != nil {
			fmt.Println("Please enter valid number!")
			return
		}

		if id == 0 {
			fmt.Println("Operation Cancelled! Number should be bigger than 0")
			return
		}

		if success := h.todoList.RemoveTodo(id); success {
			fmt.Printf("Todo with ID %d removed successfully!\n", id)
			return
		} else {
			fmt.Printf("No todo found with ID %d. Please try again.\n", id)
			continue
		}

	}

}

func (h *Handler) handlerComplete() {
	if len(h.todoList.ShowAllTodos()) == 0 {
		fmt.Println("No todos to complete.")
		return
	}

	fmt.Println("Current todos: ")
	h.handleShowTodos()

	var id int
	for {
		fmt.Println("Which todo I should mark as complete: ")
		_, err := fmt.Scan(&id)

		if err != nil {
			fmt.Println("Please enter valid number")
		}

		if id == 0 {
			fmt.Println("Operation cancelled! NUmber should be bigger than 0")
		}

		if success := h.todoList.MarkAsCompleted(id); success {
			fmt.Printf("Todo with ID %d completed!\n", id)
			return
		} else {
			fmt.Printf("No todo found with this ID %d. Please try again!\n", id)
			continue
		}
	}

}
