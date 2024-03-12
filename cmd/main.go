package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kerucko/auth/internal/app"
	"github.com/kerucko/auth/internal/config"
)

func main() {
	cfg := config.MustReadConfig()

	applicaton := app.NewApp(cfg)
	go func() {
		err := applicaton.GrpcServer.Run()
		if err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	recievedSignal := <-stop
	log.Println("stopping application due to signal: ", recievedSignal)
	applicaton.GrpcServer.Stop()
	log.Println("application stopped")
}
