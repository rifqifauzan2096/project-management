package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
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

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, &user)
	return  utils.Success(ctx, "Register user success", userResp)
}

func (uc *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Invalid request body", err.Error())
	}

	user, err := uc.userService.Login(body.Email, body.Password)

	if err != nil {
		return utils.Unauthorize(ctx, "Login Failed", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResp models.UserResponse
	_ = copier.Copy(&userResp, &user)

	return utils.Success(ctx, "Login Successful", fiber.Map{
		"access_token": token,
		"refresh_token": refreshToken,
		"user": userResp,
	})
	
}

func (uc *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := uc.userService.GetByPublicID(id)
	if err != nil {
		return utils.NotFound(ctx, "User not found", err.Error())
	}

	var userResp models.UserResponse
	err = copier.Copy(&userResp, &user)

	if err != nil {
		return utils.BadRequest(ctx, "Internal server error", err.Error())
	}

	return utils.Success(ctx, "User retrieved successfully", userResp)
}