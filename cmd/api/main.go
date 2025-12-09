package main

import (
	"log"
	"social/internal/env"
)

func main() {

	cfg := config{
		address: env.Get("SERVER_ADDRESS", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	log.Fatal(app.run(app.mount()))
}
