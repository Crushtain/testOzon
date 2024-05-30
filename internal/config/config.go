package config

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultServerAddress = "localhost:8080"

	DBhost     = "localhost"
	DBuser     = "postgres"
	DBpassword = "12345"
	DBdbname   = "postgres"
)

type Config struct {
	Host         string `env:"SERVER_ADDRESS"`
	Storage      string `env:"STORAGE_TYPE"`
	DatabasePath string `env:"DATABASE_PATH"`
}

func NewConfig() *Config {
	cfg := &Config{}
	DBConnString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DBhost, DBuser, DBpassword, DBdbname)
	flag.StringVar(&cfg.Host, "a", defaultServerAddress, "It's a Host")
	flag.StringVar(&cfg.Storage, "s", "inmemory", "Type of storage")
	flag.StringVar(&cfg.DatabasePath, "d", DBConnString, "Database conn string")

	flag.Parse()

	if envHost := os.Getenv("SERVER_ADDRESS"); envHost != "" {
		cfg.Host = envHost
	}
	if envStorage := os.Getenv("STORAGE_TYPE"); envStorage != "" {
		cfg.Host = envStorage
	}
	if envDatabse := os.Getenv("DATABASE_PATH"); envDatabse != "" {
		cfg.DatabasePath = envDatabse
	}

	return cfg
}
