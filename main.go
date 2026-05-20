package main

import (
	"log"
	"os"
	"ubigeo-service/config"
	"ubigeo-service/controller"

	"github.com/gofiber/fiber/v2"
	eureka "github.com/xuanbo/eureka-client"
)

func main() {
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           "http://localhost:8761/eureka/",
		App:                   "ubigeo-service",
		IP:                    "127.0.0.1",
		Port:                  8084,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
	})
	client.Start()
	db := config.ConnectDatabase()
	app := fiber.New()
	secretKey := os.Getenv("GATEWAY_SECRET")
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodOptions {
			return c.Next()
		}

		secret := c.Get("X-Gateway-Secret")
		if secret == "" || secret != secretKey {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Acceso denegado por el microservicio",
			})
		}
		return c.Next()
	})
	api := app.Group("/api/v1/ubigeos")
	api.Get("/search", controller.SearchUbigeo(db))
	log.Fatal(app.Listen(":8084"))
}
