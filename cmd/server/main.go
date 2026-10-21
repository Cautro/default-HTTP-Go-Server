package main

import (
	router "future-voice/cmd/http"
	config "future-voice/internal/config"
	"future-voice/internal/storage"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main()  {
	if err := godotenv.Load(); err != nil {
		slog.Info("No .env file found")
	}

	cfg, err := config.Load()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if err := os.MkdirAll("data", 0o755); err != nil {
		slog.Error("create data dir", "error", err)
		return
	}

	store, err := storage.New(cfg.DBPath)
	if err != nil {
		slog.Error("open storage", "error", err)
		return
	}
	defer store.Close()

	r := router.NewRouter(router.UseCase{
		DB: store.DB,
	})

	slog.Info("Server start", "addr", cfg.Port)
	if err := r.Run(cfg.DBPath); err != nil {
		slog.Error("Server failed", "err", err)
		return
	}
}