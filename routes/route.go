package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/rifqifauzan2096/project-management/config"
	"github.com/rifqifauzan2096/project-management/controllers"
	"github.com/rifqifauzan2096/project-management/utils"
)

func Setup(app *fiber.App, uc *controllers.UserController) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	app.Post("/v1/auth/register", uc.Register)
	app.Post("/v1/auth/login", uc.Login)
	
	//JWT Protected routes can be added here
	api := app.Group("/api/v1", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.AppConfig.JWTSecret),
		ContextKey: "user",
		ErrorHandler: func (c *fiber.Ctx, err error) error {
			return utils.Unauthorize(c, "Error Unauthorized", err.Error())			
		},
	}))

	//endpoint protected can be added here
	userGroup := api.Group("/users")
	userGroup.Get("/:id", uc.GetUser)
}