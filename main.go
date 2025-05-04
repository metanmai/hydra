package main

import (
	"hydra/models"
	"hydra/routes"
)

func main() {
	r := routes.SetupRouter()
	models.SetupDatabase()
	r.Run("0.0.0.0:8080")
}
