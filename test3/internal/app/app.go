package app

import (
	"github.com/dingowd/CyberZilla/test3/internal/logger"
	"github.com/dingowd/CyberZilla/test3/internal/pusher"
	"github.com/dingowd/CyberZilla/test3/internal/storage"
)

type App struct {
	Logg    logger.Logger
	Storage storage.Storage
	Pusher  pusher.Pusher
}

func New(logger logger.Logger, storage storage.Storage, push pusher.Pusher) *App {
	return &App{
		Logg:    logger,
		Storage: storage,
		Pusher:  push,
	}
}
