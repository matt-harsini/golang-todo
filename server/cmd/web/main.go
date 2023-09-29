package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
}

func main() {
	
	app := &application{}
	server := &http.Server{
		Handler: app.routes(),
	}
	server.ListenAndServe()
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "golang:pass@/golangtodo?parseTime=true")

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
