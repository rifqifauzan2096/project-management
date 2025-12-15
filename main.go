package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rifqifauzan2096/project-management/config"
	"github.com/rifqifauzan2096/project-management/controllers"
	"github.com/rifqifauzan2096/project-management/database/seed"
	"github.com/rifqifauzan2096/project-management/repositories"
	"github.com/rifqifauzan2096/project-management/routes"
	"github.com/rifqifauzan2096/project-management/services"
)


func main() {
	//entry point of the application
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()

	app := fiber.New()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server is running on port : ", port)

	log.Fatal(app.Listen(":" + port))	
}	