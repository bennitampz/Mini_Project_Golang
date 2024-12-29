package handlers

import (
	"net/http"

	"tokoku_go/config"
	"tokoku_go/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetProfile(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var user models.User
	if result := config.DB.First(&user, "id = ?", userID); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	var store models.Store
	if result := config.DB.Where("user_id = ?", userID).First(&store); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Store not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user":  user,
		"store": store,
	})
}
