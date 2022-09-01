//Filename:  cmd/api/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Application Version number
const version = "1.0.0"

//Configuration Settings

type config struct {
	port int
	env  string //development, staging, production
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config
	//read in the flafs that are need to populate our config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development | staging | production)")
	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//create an instance of our application struct

	app := &application{
		config: cfg,
		logger: logger,
	}

	//create our new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	//create http server

	srv := &http.Server{

		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	//start server

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
