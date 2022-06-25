package book

import (
	"github.com/gofiber/fiber/v2"
)

//Get all books
func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All books listed")
}

//Get a specific book
func GetBook(c *fiber.Ctx) error {
	return c.SendString("Single Book")
}

//Add new book
func NewBook(c *fiber.Ctx) error {
	return c.SendString("New Book")
}

//Delete a specific book
func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete Book")
}
