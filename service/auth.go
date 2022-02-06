package service

import (
	"log"	
	
	"new-proj/dto"
	entity "new-proj/entities"
	"new-proj/helper"
	"new-proj/repositories"

	"github.com/mashingan/smapping"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByMail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToBeCreated := entity.User{}

	err := smapping.FillStruct(&userToBeCreated, smapping.MapFields(&user))

	if err != nil {
		log.Fatalf("failed to map %v", err)
	}

	res := s.userRepo.InsertUser(userToBeCreated)

	return res
}

func (s *authService) FindByMail(email string) entity.User {
	return s.userRepo.FindByEmail(email)
}

func (s *authService) IsDuplicateEmail(email string) bool {
	res := s.userRepo.IsDuplicateEmail(email)

	return !(res.Error == nil)
}

func (s *authService) VerifyCredential(email string, password string) interface{} {
	res := s.userRepo.VerifyCredential(email, password)

	if v, ok := res.(entity.User); ok {
		comparedPassword := helper.ComparePassword(v.Password, [] byte(password))

		if v.Email == email && comparedPassword {
			return res
		}

		return false
	}

	return false
}