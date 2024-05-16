package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"pos-go-redis-limiter/application"
	"pos-go-redis-limiter/infrastructure"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config, err := infrastructure.Load()
	if err != nil {
		log.Fatal("Error loading mapper configs")
	}

	r := application.StartupApp(config)
	fmt.Println("Server is running on port 8080...")
	_ = r.Run(":8080")
}
