package main

import (
	"fmt"
	"net/http"
)

func handleTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Hello world")
}

func main() {
	fmt.Printf("Hello world")
	mux := http.NewServeMux()

	mux.HandleFunc("/submit", handleTodo)


	http.ListenAndServe(":4000", mux)
}
