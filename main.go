package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/malma/malma-blog-be/controller"
	"github.com/malma/malma-blog-be/mapper"
	"github.com/malma/malma-blog-be/pkg/storage"
	"github.com/malma/malma-blog-be/repository"
	"github.com/malma/malma-blog-be/validator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(config.Database))
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// pkg
	storageManager := storage.NewStorageManager("statics")

	// blog
	blogValidator := validator.NewBlogValidator()
	blogRepository := repository.NewBlogRepository(db)
	blogMapper := mapper.NewBlogMapper()
	blogController := controller.NewBlogController(blogValidator, blogRepository, blogMapper, storageManager)

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		ErrorHandler:  controller.ErrorController(),
	})
	app.Static("/", "./statics")

	app.Route("/blogs", blogController.Routing)

	if err := app.Listen(fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port)); err != nil {
		panic(err)
	}
}
