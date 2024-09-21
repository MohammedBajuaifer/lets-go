package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})

	logger := slog.New(loggerHandler)

	app := &application{logger: logger}

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger.Info("Starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())

	os.Exit(1)
}
