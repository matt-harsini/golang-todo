package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"todo/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	todos    *models.TodosModel
}

func main() {
	addr := flag.String("addr", ":4000", "Port number")
	dsn := flag.String("dsn", "golang:pass@/golangtodo?parseTime=true", "MySQL connection string")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		todos:    &models.TodosModel{DB: db},
	}

	server := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	server.ListenAndServe()
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "golang:pass@/golangtodo?parseTime=true")

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
