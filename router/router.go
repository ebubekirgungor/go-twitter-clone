package router

import (
	"go-twitter-clone/controllers"
	"go-twitter-clone/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Post("/check_user", controllers.CheckUser)

	users := api.Group("/users")
	users.Get("/:id", controllers.User)
	users.Post("/", controllers.AddUser)
	users.Patch("/:id", middleware.Protected(), controllers.UpdateUser)
	users.Delete("/:id", middleware.Protected(), controllers.DeleteUser)

	tweets := api.Group("/tweets")
	tweets.Get("/", controllers.AllTweets)
	tweets.Get("/:id", controllers.Tweet)
	tweets.Post("/", middleware.Protected(), controllers.AddTweet)
	tweets.Delete("/:id", middleware.Protected(), controllers.DeleteTweet)
}
