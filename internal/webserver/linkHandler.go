package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0c1337/linkshortener/internal/models"
)

func (ws *WebServer) CreateLink(c *fiber.Ctx) error {
	link := new(models.Link)
	if err := c.BodyParser(link); err != nil {
		return c.SendStatus(501)
	}

	var newLink models.Link
	newLink.Discriminator = link.Discriminator
	newLink.Url = link.Url

	response := ws.db.Create(&newLink)
	if response.Error != nil {
		return response.Error
	}

	return c.SendStatus(200)
}

func (ws *WebServer) GetLinkByDiscriminator(c *fiber.Ctx) error {
	var link models.Link

	response := ws.db.Where("discriminator = ?", c.Params("discriminator")).First(&link)
	if response.Error != nil {
		return response.Error
	}

	return c.Status(200).JSON(link)
}

func (ws *WebServer) GetLinkById(c *fiber.Ctx) error {
	var link models.Link

	response := ws.db.First(&link, c.Params("id"))
	if response.Error != nil {
		return response.Error
	}

	return c.Status(200).JSON(link)
}

func (ws *WebServer) DeleteLink(c *fiber.Ctx) error {
	var link models.Link

	response := ws.db.First(&link, c.Params("id")).Delete(&link)
	if response.Error != nil {
		return response.Error
	}

	return c.Status(200).JSON(fiber.Map{
		"Deleted": "true",
	})
}

func (ws *WebServer) redirect(c *fiber.Ctx) error {
	var link models.Link

	response := ws.db.Where("discriminator = ?", c.Params("redirect")).First(&link)
	if response.Error != nil {
		return response.Error
	}

	return c.Redirect(link.Url, fiber.StatusTemporaryRedirect)
}
