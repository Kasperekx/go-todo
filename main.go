package main

import (
	"github.com/kasperekx/todo/internal/cli"
	"github.com/kasperekx/todo/internal/storage"
)

func main() {
	todoList := storage.NewToDoList()
	handler := cli.NewHandler(todoList)
	handler.Run()
}
