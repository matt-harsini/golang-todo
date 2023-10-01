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

func (model *TodosModel) Insert(todo string) (int64, error) {
	stmt := `INSERT INTO todos (todo, created) VALUES(?, UTC_TIMESTAMP())`
	result, err := model.DB.Exec(stmt, todo)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
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

func (model *TodosModel) Delete(id string) error {
	stmt := `DELETE from todos WHERE id = ?`

	result, err := model.DB.Exec(stmt, id)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func (model *TodosModel) Edit(id string) error {
	stmt := `DELETE from todos WHERE id = ?`

	result, err := model.DB.Exec(stmt, id)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}
