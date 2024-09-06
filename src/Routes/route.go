package Routes

import (
	"book-fiber/Controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/books", Controller.GetBook)
	app.Post("/books", Controller.CreateBook)
	app.Get("/books/search", Controller.GetBooksByAuthor)
}
