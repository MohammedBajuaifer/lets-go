package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	logger *slog.Logger
}

const createTable string = `
	CREATE TABLE IF NOT EXISTS snippets (
	id INTEGER NOT NULL PRIMARY KEY,
	time DATETIME NOT NULL,
	description TEXT
	)
`

func main() {

	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		slog.Error(err.Error())
	}

	_, err = db.Exec(createTable)
	if err != nil {
		slog.Error(err.Error())
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})

	logger := slog.New(loggerHandler)

	app := &application{logger: logger}

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger.Info("Starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())

	os.Exit(1)
}
