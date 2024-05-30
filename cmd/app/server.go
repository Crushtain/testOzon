package main

import (
	"net/http"

	"github.com/Crushtain/testOzon/internal/app"
	"github.com/Crushtain/testOzon/internal/config"
	"github.com/Crushtain/testOzon/internal/route"
)

func main() {
	cfg := config.NewConfig()
	App := app.NewApp(cfg)
	App.ConfigureStorage(cfg)
	router := route.Route(App)

	err := http.ListenAndServe(cfg.Host, router)
	if err != nil {
		panic(err)
	}
}
