package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 8040, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|port)")
	flag.Parse()

	app := &application{
		config: cfg,
		logger: log.New(log.Writer(), "API: ", log.Ldate|log.Ltime),
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)

	log.Printf("starting %s server on %s", cfg.env, addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}