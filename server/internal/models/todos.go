package models

import (
	"database/sql"
	"time"
)

type Todos struct {
	ID      int
	Todo    string
	Created time.Time
}

type TodosModel struct {
	DB *sql.DB
}

func (model *TodosModel) Insert(todo string) (int, error) {
	return 0, nil
}
