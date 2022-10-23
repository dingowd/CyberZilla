package app

import (
	"github.com/dingowd/CyberZilla/test3/internal/logger"
	"github.com/dingowd/CyberZilla/test3/internal/storage"
)

type App struct {
	Logg    logger.Logger
	Storage storage.Storage
}

func New(logger logger.Logger, storage storage.Storage) *App {
	return &App{
		Logg:    logger,
		Storage: storage,
	}
}
