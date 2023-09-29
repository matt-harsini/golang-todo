package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
)

func (app *application) handleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method != "POST" {
		return
	}
	// stmt := `INSERT INTO todos (todo, created) VALUES(?, UTC_TIMESTAMP())`
	var todo []string
	// err := json.NewDecoder(r.Body).Decode(&todo)
	fmt.Print(todo, len(todo))
	// db := openDB()
	// _, err = db.Exec(stmt, todo[0])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
