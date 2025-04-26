package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	db, err := sqlx.Open("postgres", "user=admin dbname=contact_management sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	api := app.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Println("Listening on :" + port)
	log.Fatal(app.Listen(":" + port))
}
