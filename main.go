package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyyyyyyyyyyyyyy/golang-gofiber/controllers/bookcontroller"
	"github.com/kyyyyyyyyyyyyyy/golang-gofiber/models"
)

func main() {

	models.ConnectDatabase()

	app := fiber.New()

	//grouping
	api := app.Group("/api")
	book := api.Group("/book")

	book.Get("", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("", bookcontroller.Store)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":8000")

}
