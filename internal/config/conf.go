package config

import (
	"os"
)

type Config struct {
	Port string
	DBPath string
}

func Load() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		panic("Port in env don't exist")
	}

	dbPath := os.Getenv("DBPATH")
	if dbPath == "" {
		panic("DB path in env don't exist")
	}

	return Config{
		Port: port,
		DBPath: dbPath,
	}, nil
}