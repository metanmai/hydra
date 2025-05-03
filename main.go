package main

import (
	"github.com/henvo/golang-gin-gorm-starter/models"
	"github.com/henvo/golang-gin-gorm-starter/routes"
)

func main() {
	r := routes.SetupRouter()
	models.SetupDatabase()
	r.Run()
}
