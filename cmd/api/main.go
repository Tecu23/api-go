package main

import (
	"flag"
	"os"
	"sync"

	"github.com/Tecu23/api-go/internal/data"
	"github.com/Tecu23/api-go/internal/jsonlog"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
	wg     sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.Parse()

	// Initialize a new logger which writes to std out, prefixed with curr date & tim
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &application{
		config: cfg,
		logger: logger,
		// models: data.NewModels(db),
	}

	err := app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
