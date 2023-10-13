package tweet

import (
	"go-twitter-clone/database"
	"go-twitter-clone/models"

	"github.com/gofiber/fiber/v2"
)

func AllTweets(c *fiber.Ctx) error {
	tweets := []models.Tweet{}
	database.DB.Db.Preload("User").Find(&tweets)

	var all_tweets []fiber.Map
	for _, tweet := range tweets {
		all_tweets = append(all_tweets, fiber.Map{
			"id":      tweet.ID,
			"content": tweet.Content,
			"user": fiber.Map{
				"id":       tweet.User.ID,
				"username": tweet.User.Username,
				"email":    tweet.User.Email,
			},
		})
	}
	return c.Status(200).JSON(all_tweets)
}

func Tweet(c *fiber.Ctx) error {
	tweet := models.Tweet{}
	database.DB.Db.Preload("User").First(&tweet, c.Params("id"))
	return c.Status(200).JSON(fiber.Map{
		"id":      tweet.ID,
		"content": tweet.Content,
		"user": fiber.Map{
			"id":       tweet.User.ID,
			"username": tweet.User.Username,
			"email":    tweet.User.Email,
		},
	})
}

func AddTweet(c *fiber.Ctx) error {
	tweet := new(models.Tweet)
	if err := c.BodyParser(tweet); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&tweet)

	return c.Status(200).JSON(tweet)
}

func Delete(c *fiber.Ctx) error {
	tweet := []models.Tweet{}
	database.DB.Db.Where("id = ?", c.Params("id")).Delete(&tweet)

	return c.Status(200).JSON("deleted")
}
