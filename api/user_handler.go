package api

import (
	"github.com/Jieqstu/go-hotel-reservation/api/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName: "Petter",
	}
	return c.JSON(u)
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON("James")
}