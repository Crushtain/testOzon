package app

import (
	"github.com/Crushtain/testOzon/internal/config"
	"github.com/Crushtain/testOzon/internal/database"
	"github.com/Crushtain/testOzon/internal/storage"
)

type App struct {
	Config   *config.Config
	Storage  storage.Storage
	Database *database.DB
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Config:   cfg,
		Database: database.Init(cfg.DatabasePath),
	}
}

func (a *App) ConfigureStorage(cfg *config.Config) {
	switch cfg.Storage {
	case "postgres":
		a.Storage = storage.NewPostgresStorage(a.Database)
	default:
		a.Storage = storage.NewInMemory()
	}
}
