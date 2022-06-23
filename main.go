package main

import (
	"fmt"

	//"mysql-books/book"
	"github.com/SubratBarik/mysql-books/book"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	//app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	fmt.Println("Books Empire...")

	app := fiber.New()

	setupRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Books Empire")
	// })

	app.Listen(":3000")
}
