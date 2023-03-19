package main

import (
	"log"

	"github.com/co-codin/go-grpc-template/internal/db"
	"github.com/co-codin/go-grpc-template/internal/rocket"
	"github.com/joho/godotenv"
)

func Run() error {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	rocketStore, err := db.New()

	err = rocketStore.Migrate()

	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	if err != nil {
		return err
	}

	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
