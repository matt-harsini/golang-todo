package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/submit", app.submitTodo)
	mux.HandleFunc("/get-all", app.getAllTodos)
	mux.HandleFunc("/delete", app.deleteTodo)
	mux.HandleFunc("/edit", app.editTodo)
	return mux
}
