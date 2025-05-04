package main

import (
	"hydra/models"
	"hydra/routes"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.Println("Initializing application...")

	log.SetFormatter(&log.TextFormatter{
        ForceColors:   true,
        FullTimestamp: true,
    })
	log.SetLevel(log.InfoLevel)
	
	log.WithFields(map[string]interface{}{
		"app": "hydra",
		"version": "1.0.0",
	}).Info("Starting application...")

	// Initialize database
	models.SetupDatabase()
	log.WithFields(map[string]interface{}{
		"db": "hydra_db",
		"user": "metanmai",
		"host": "localhost",
		"port": "5432",
	}).Info("Database connection established")
	
	// Initialize routes
	
}


func main() {
	r := routes.SetupRouter()
	log.WithFields(map[string]interface{}{
		"routes": "api",
	}).Info("Routes initialized")

	r.Run("0.0.0.0:8080")
}
