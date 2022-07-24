package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"user_api/component"
	"user_api/server"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	appContext, err := component.NewAppContext(ctx)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	s := server.NewServer(appContext)

	go func() {
		_ = <-c
		cancel()
		if err = appContext.Close(); err != nil {
			log.Fatalf("%s", err.Error())
		}
		s.Shutdown()
	}()

	if err = s.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
