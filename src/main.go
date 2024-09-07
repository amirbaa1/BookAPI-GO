package main

import (
	"book-fiber/Config"
	"book-fiber/Model"
	"book-fiber/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func ConnectDatabase() {
	Config.Connect()
	log.Println("Connected to database")
	db := Config.GetDB()

	err := db.AutoMigrate(&Model.Book{}, &Model.Author{}, &Model.Auth{})
	if err != nil {
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}
	log.Println("Database migrated")
}

func main() {
	ConnectDatabase()

	app := fiber.New()
	//jwt := Middlewares.AuthMiddleware("SECRET:)")

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World TEST Fiber BookAPI!")
	})
	Routes.SetupRoutes(app)
	Routes.AuthRoutes(app)

	log.Fatal(app.Listen(":3000"))

	//app.Listen(":3000")
}
