package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) submitTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.Write([]byte("Ok"))
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var todo string
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		app.errorLog.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	id, err := app.todos.Insert(todo)

	if err != nil {
		app.errorLog.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(id)
}

func (app *application) getAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.Write([]byte("Ok"))
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	payload, err := app.todos.GetAll()

	if err != nil {
		app.errorLog.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(payload)
}

func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if r.Method == http.MethodOptions {
		w.Write([]byte("Ok"))
		return
	}

	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := app.todos.Delete(r.URL.Query().Get("id"))

	if err != nil {
		app.errorLog.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Write([]byte("Ok"))
}

func (app *application) editTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if r.Method == http.MethodOptions {
		w.Write([]byte("Ok"))
		return
	}

	if r.Method != http.MethodPatch {
		w.Header().Set("Allow", http.MethodPatch)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var todo string
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		app.errorLog.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = app.todos.Edit(r.URL.Query().Get("id"), todo)

	if err != nil {
		app.errorLog.Print(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Write([]byte("Ok"))
}
