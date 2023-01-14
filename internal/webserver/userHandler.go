package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0c1337/linkshortener/internal/models"
)

func (ws *WebServer) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(501)
	}

	var newUser models.User
	newUser.Username = user.Username
	newUser.PasswordHash = ws.auth.CreateHash(user.PasswordHash)

	ws.db.Create(&newUser)

	return c.SendStatus(200)
}

func (ws *WebServer) GetUser(c *fiber.Ctx) error {
	var user models.User

	ws.db.First(&user, c.Params("id"))
	return c.Status(200).JSON(user)
}

func (ws *WebServer) DeleteUser(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
