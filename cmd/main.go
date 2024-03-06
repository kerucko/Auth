package main

import (
	"github.com/kerucko/auth/internal/app"
	"github.com/kerucko/auth/internal/config"
)

func main() {
	cfg := config.MustReadConfig()

	applicaton := app.NewApp(cfg)
	err := applicaton.GrpcServer.Run()
	if err != nil {
		panic(err)
	}
}
