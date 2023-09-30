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
	stmt := `INSERT INTO todos (todo, created) VALUES(?, UTC_TIMESTAMP())`
	_, err := model.DB.Exec(stmt, todo)

	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (model *TodosModel) GetAll() ([]Todos, error) {
	stmt := `SELECT * FROM todos`
	rows, err := model.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	payload := []Todos{}

	for rows.Next() {
		var todo Todos
		err = rows.Scan(&todo.ID, &todo.Todo, &todo.Created)
		if err != nil {
			return nil, err
		}
		payload = append(payload, todo)
	}

	return payload, nil
}
