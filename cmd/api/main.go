package main

import (
	"fmt"
	"log"
	db "social/internal/database"
	postRepository "social/internal/posts/repository"
	postService "social/internal/posts/service"
	postTransport "social/internal/posts/transport"
	userRepository "social/internal/users/repository"
	userService "social/internal/users/service"
	userTransport "social/internal/users/transport"
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
		userHandler: userTransport.NewUserHandler(
			userService.NewUserService(
				userRepository.NewPostgres(db),
			),
		),
		postHandler: postTransport.NewPostHandler(
			postService.NewPostService(
				postRepository.NewPostgres(db),
			),
		),
	}

	app := &application{
		config:  cfg,
		handler: handler,
	}

	log.Fatal(app.run(app.mount()))
}
