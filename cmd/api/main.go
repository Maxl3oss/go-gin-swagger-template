package main

import (
	"log"
	"role-management/internal/api/routes"
	"role-management/internal/config"
	"role-management/pkg/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// set db
	db, err := database.PostgreSQLConnection()
	if err != nil {
		panic("failed to connect database")
	}

	r := routes.SetupRouter(cfg, db)
	r.Run(":" + cfg.Port)
}
