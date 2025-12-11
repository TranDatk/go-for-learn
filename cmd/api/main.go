package main

import (
	"log"
	"social/internal/users/repository"
	"social/internal/users/service"
	"social/internal/users/transport"
	"social/pkg/env"
)

func main() {
	cfg := config{
		address: env.Get("SERVER_ADDRESS", ":8080"),
	}

	handler := handler{
		userHandler: transport.NewHandler(
			service.New(
				repository.NewMemory(),
			),
		),
	}

	app := &application{
		config:  cfg,
		handler: handler,
	}

	log.Fatal(app.run(app.mount()))
}
