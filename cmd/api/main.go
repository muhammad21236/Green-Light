package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	Port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	stdLogger := log.New(os.Stdout, "", log.LstdFlags)

	app := application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.Port),
		Handler:     mux,
		IdleTimeout: 60 * time.Second,
		ReadTimeout: 5 * time.Second,
		ErrorLog:    stdLogger,
	}

	logger.Info("Starting server", slog.Int("port", cfg.Port))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Error starting server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
