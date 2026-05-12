package main

import (
	"log"
	"ubigeo-service/config"
	"ubigeo-service/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.ConnectDatabase()
	app := fiber.New()
	api := app.Group("/api/v1/ubigeos")
	api.Get("/search", controller.SearchUbigeo(db))
	log.Fatal(app.Listen(":8084"))
}
