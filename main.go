package main

import (
	"log"

	"github.com/alisinasoltani/goFiber/database"
	"github.com/alisinasoltani/goFiber/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

	app.Post("/api/orders", routes.CreateOrder)
	// app.Get("/api/orders", routes.GetProducts)
	// app.Get("/api/orders/:id", routes.GetProduct)
	// app.Put("/api/orders/:id", routes.UpdateProduct)
	// app.Delete("/api/orders/:id", routes.DeleteProduct)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
