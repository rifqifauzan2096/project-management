package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/rifqifauzan2096/project-management/models"
	"github.com/rifqifauzan2096/project-management/repositories"
	"github.com/rifqifauzan2096/project-management/utils"
)

type UserService interface{
	Register(user *models.User) error
}

type userService struct{
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService{
	return &userService{repo}
}

func (s *userService) Register(user *models.User) error {
	//cek email sudah terdaftar
	//membuat hashing password
	//set role
	//simpat user ke DB

	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID !=0 {
		return errors.New("email already registered")
	}

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed
	user.Role = "user"
	user.PublicID = uuid.New()

	return s.repo.Create(user)

}