package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func handleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	var todo []string 
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(todo)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/submit", handleTodo)

	http.ListenAndServe(":4000", mux)
}

func openDB() *sql.DB {
	db, err := sql.Open("mysql", "golang:pass@/golangtodo")

	if err != nil {
		panic(err)
	}

	return db
}
