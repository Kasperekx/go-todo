package models

import (
	"fmt"
	"time"
)

type Todo struct {
	ID        int
	Name      string
	Completed bool
	CreatedAt time.Time
}

func (t Todo) FormattedDate() string {
	return t.CreatedAt.Format("2006-01-02 15:04:05")
}

func (t Todo) String() string {
	return fmt.Sprintf("Todo: %s (Created: %s, Completed: %v)", t.Name, t.FormattedDate(), t.Completed)
}
