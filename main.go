package main

import (
	"log"
	"ubigeo-service/config"
	"ubigeo-service/controller"

	"github.com/gofiber/fiber/v2"
	eureka "github.com/xuanbo/eureka-client"
)

func main() {
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://localhost:8761/eureka/",
		App:                   "ubigeo-service",
		Port:                  8084,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
	})
	client.Start()
	db := config.ConnectDatabase()
	app := fiber.New()
	api := app.Group("/api/v1/ubigeos")
	api.Get("/search", controller.SearchUbigeo(db))
	log.Fatal(app.Listen(":8084"))
}
