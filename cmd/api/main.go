package main

import (
	"log"

	"github.com/Wizardsmile1412/near-changs-backend/internal/config"
	"github.com/Wizardsmile1412/near-changs-backend/internal/database"
	"github.com/Wizardsmile1412/near-changs-backend/internal/handler"
)

func main() {
	//Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	//Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")

	//Setup router
	router := handler.SetupRouter(db)

	//Start server
	log.Printf("Near Changs API starting on port %s...", cfg.ServerPort)
	err = router.Run(":" + cfg.ServerPort)
	if err != nil {
		log.Fatalf("failed to start server:%v", err)
	}
}
