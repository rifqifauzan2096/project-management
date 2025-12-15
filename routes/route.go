package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rifqifauzan2096/project-management/controllers"
)

func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app.Post("/v1/auth/register", uc.Register)
}