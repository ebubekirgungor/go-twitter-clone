package database

import (
	"fmt"
	"log"

	"go-twitter-clone/config"
	"go-twitter-clone/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=localhost user=postgres password=%s dbname=%s port=%s sslmode=disable",
		config.Config("DATABASE_PASSWORD"),
		config.Config("DATABASE_NAME"),
		config.Config(("DATABASE_PORT")))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Following{})
	db.AutoMigrate(&models.Follower{})
	db.AutoMigrate(&models.Tweet{})

	DB = Dbinstance{
		Db: db,
	}
}
