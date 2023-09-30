package main

import (
	// "encoding/json"
	"encoding/json"
	"net/http"
)

func (app *application) submitTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	var todo string
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		app.errorLog.Fatal(err)
	}

	app.todos.Insert(todo)

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
