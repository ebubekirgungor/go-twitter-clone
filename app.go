package main

import (
	"go-twitter-clone/controllers/tweet"
	"go-twitter-clone/controllers/user"
	"go-twitter-clone/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/api/users", user.AllUsers)
	app.Get("/api/users/:id", user.User)
	app.Post("/api/users", user.AddUser)
	app.Delete("/api/users/:id", user.Delete)

	app.Get("/api/tweets", tweet.AllTweets)
	app.Get("/api/tweets/:id", tweet.Tweet)
	app.Post("/api/tweets", tweet.AddTweet)
	app.Delete("/api/tweets/:id", tweet.Delete)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Static("/", "./client/dist")

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
