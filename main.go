package main

import (
	"fmt"
	"log"
	"pos-go-redis-limiter/application"
	"pos-go-redis-limiter/infrastructure"
)

func main() {
	config, err := infrastructure.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := application.StartupApp(config)
	fmt.Println("Server is running on port 8080...")
	_ = r.Run(":8080")
}
