package controller

import (
	"ubigeo-service/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SearchUbigeo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		state := c.Query("state")
		province := c.Query("province")
		county := c.Query("county")
		var results []model.Ubigeo
		query := db.Model(&model.Ubigeo{})
		if state != "" {
			query = query.Where("departamento = ?", state)
		}
		if province != "" {
			query = query.Where("provincia = ?", province)
		}
		if county != "" {
			query = query.Where("distrito = ?", county)
		}
		result := query.Order("departamento, provincia, distrito").Find(&results)
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error occurred during database search"})
		}
		return c.Status(200).JSON(fiber.Map{
			"status": true,
			"data":   results,
		})
	}

}
