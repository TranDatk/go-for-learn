package main

import (
	"fmt"
	"log"
	db "social/internal/database"
	"social/internal/users/repository"
	"social/internal/users/service"
	"social/internal/users/transport"
	"social/pkg/env"
)

func main() {
	cfg := config{
		addr: env.Get("SERVER_ADDRESS", ":8080"),
		db: dbConfig{
			addr:         env.Get("DB_ADDRESS", ""),
			maxOpenConns: env.Get("DB_MAX_CONNS", 30),
			maxIdleConns: env.Get("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.Get("DB_MAX_IDLE_TIME", "5m"),
		},
	}

	db, err := db.New(
		"postgres",
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database connection pool established")

	handler := handler{
		userHandler: transport.NewUserHandler(
			service.NewUserService(
				repository.NewPostgres(db),
			),
		),
	}

	app := &application{
		config:  cfg,
		handler: handler,
	}

	log.Fatal(app.run(app.mount()))
}
