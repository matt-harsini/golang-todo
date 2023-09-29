package main

import "net/http"

type application struct {
}

func main() {
	app := &application{}
	server := &http.Server{
		Handler: app.routes(),
	}
	server.ListenAndServe()
}
