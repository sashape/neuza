package model

import (
	"neuza/backend/database"
	"neuza/backend/responses"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetAll(model interface{}, database *database.Database, c *fiber.Ctx) error {
	db = database.DB

	db.AutoMigrate(model)
	result := db.Find(model)

	if result.Error != nil {
		switch err := result.Error; err {
		case gorm.ErrRecordNotFound:
			return responses.NotFound(c, "Items not found")
		default:
			return responses.InternalServerError(c, "Server error")
		}
	}
	return responses.OK(c, fiber.Map{"items": model})
}
