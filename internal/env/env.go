package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type StrInt interface {
	~string | ~int
}

func Get[T StrInt](key string, fallback T) T {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	var zero T

	switch any(zero).(type) {
	case int:
		i, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		}
		return any(i).(T)

	case string:
		return any(val).(T)
	}

	return fallback
}
