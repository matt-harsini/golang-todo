package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/submit", app.submitTodo)
	mux.HandleFunc("/get-all", app.getAllTodos)
	return mux
}
