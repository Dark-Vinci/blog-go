package service

import (
	"fmt"
	
	"new-proj/dto"
	entity "new-proj/entities"
	"new-proj/repositories"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(UserId string) entity.User
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (c *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))

	if err != nil {
		fmt.Printf("failed to map %v", err)
	}

	updatedUser := c.userRepo.UpdateUser(userToUpdate)

	return updatedUser
}

func (c *userService) Profile(UserId string) entity.User {
	return c.userRepo.ProfileUser(UserId)
}