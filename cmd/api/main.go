package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a map which holds the information that we want to send in the response.
	data := map[string]string{
		"status": "available", "environment": app.config.env, "version": version,
	}

	// Pass the map to the json.Marshal() function. This returns a []byte slice
	// containing the encoded JSON. If there was an error, we log it and send the client
	// a generic error message.
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// Append a newline to the JSON. This is just a small nicety to make it easier to
	// view in terminal applications.
	js = append(js, '\n')

	// At this point we know that encoding the data worked without any problems, so we
	// can safely set any necessary HTTP headers for a successful response.
	// Set the "Content-Type: application/json" header on the response.
	// If you forget to do this, Go will default to sending a "Content-Type: text/plain; charset=utf-8"
	// header instead.
	w.Header().Set("Content-Type", "application/json")

	// Use w.Write() to send the []byte slice containing the JSON as the response body.
	w.Write(js)
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	// Example:  go run ./cmd/api -port=3030 -env=production
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)
}
