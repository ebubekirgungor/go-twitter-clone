package user

import (
	"go-twitter-clone/database"
	"go-twitter-clone/models"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Preload("Tweets").Find(&users)

	return c.Status(200).JSON(users)
}

func User(c *fiber.Ctx) error {
	user := models.User{}
	database.DB.Db.Preload("Tweets").First(&user, c.Params("id"))

	return c.Status(200).JSON(user)
}

func AddUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func Delete(c *fiber.Ctx) error {
	user := []models.User{}
	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&user)

	return c.Status(200).JSON("deleted")
}
