package router

import (
	"go-twitter-clone/controllers/tweet"
	"go-twitter-clone/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	users := api.Group("/users")
	users.Get("/", user.AllUsers)
	users.Get("/:id", user.User)
	users.Post("/", user.AddUser)
	users.Delete("/:id", user.Delete)

	tweets := api.Group("/tweets")
	tweets.Get("/", tweet.AllTweets)
	tweets.Get("/:id", tweet.Tweet)
	tweets.Post("/", tweet.AddTweet)
	tweets.Delete("/:id", tweet.Delete)
}
