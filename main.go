package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	// "172.17.11.15/12500027/todo-backend/config"
	// "172.17.11.15/12500027/todo-backend/db"
	// "172.17.11.15/12500027/todo-backend/routes"
	"todo-backend/config"
	"todo-backend/db"
	"todo-backend/routes"
)

func main() {
	// load .env (optional — if no .env, env vars used)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found (continuing with system env vars)")
	}

	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// connect to DB
	gormDB, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	r := routes.Setup(gormDB)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
