package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"url_rest_api/internal/config"
	"url_rest_api/internal/logger/sl"
	"url_rest_api/internal/storage/postgre"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoadConfig()

	log := setupLogger(cfg.Env)

	log.Info("starting", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	if err := godotenv.Load(); err != nil {
		log.Error("Error loading .env file")
	}
	db, err := postgre.InitPostgre(postgre.ConfigDB{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		User:     cfg.DB.User,
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   cfg.DB.DBName,
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = db
	// TODO:init router: chi, render

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
