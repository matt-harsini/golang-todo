package models

import (
	"database/sql"
	"fmt"
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
	stmt := `INSERT INTO todos (todo, created) VALUES(?, UTC_TIMESTAMP())`
	_, err := model.DB.Exec(stmt, todo)

	if err != nil {
		fmt.Println(err)
	}

	return 0, nil
}
