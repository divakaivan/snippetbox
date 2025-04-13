package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// holds app-wide deps
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// log.New returns a concurrency-safe logger and we can share the same logger to many goroutines
	// but if multiple loggers write to the same dst we need to be careful
	// it's better to log to stdout and redirect to a file at runtime
	// alternatively, we can open a file in Go and use it as a log dst
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// init a new instance containing the deps
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// setup http server to use the custom errorLog
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
