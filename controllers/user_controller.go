package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifqifauzan2096/project-management/models"
	"github.com/rifqifauzan2096/project-management/services"
	"github.com/rifqifauzan2096/project-management/utils"
)


type UserController struct{
	userService services.UserService
}

func NewUserController(s services.UserService) *UserController{
	return &UserController{userService: s}
}

func (uc *UserController) Register(ctx *fiber.Ctx) error {
	//handler register user

	user := new(models.User)
	if err  := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	if err := uc.userService.Register(user); err != nil {
		return utils.BadRequest(ctx, "Failed to register user", err.Error())
	}

	return  utils.Success(ctx, "Register user success", user)
}