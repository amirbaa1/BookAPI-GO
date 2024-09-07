package Routes

import (
	"book-fiber/Controller"
	"book-fiber/Middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	jwtMiddleware := Middlewares.AuthMiddleware("SECRET:)")

	app.Get("/books", Controller.GetBook)
	app.Post("/books", jwtMiddleware, Controller.CreateBook)
	app.Get("/books/search", Controller.GetBooksByAuthor)
	app.Get("/books/:id", Controller.GetBookById)
	app.Get("/books/:title", Controller.GetBookByTitle)
	app.Get("/books/:publisher", Controller.GetBookByPublisher)
	app.Get("/authors", Controller.GetAuthorList)
}

func AuthRoutes(app *fiber.App) {
	app.Post("/auth/register", Controller.Register)
	app.Post("/auth/login", Controller.Login)

	jwtMiddleware := Middlewares.AuthMiddleware("SECRET:)")
	app.Get("/auth/profile", jwtMiddleware, Controller.Profile)
}
