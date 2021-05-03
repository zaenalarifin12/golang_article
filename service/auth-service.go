package service

import (
	"github.com/mashingan/smapping"
	"github.com/zaenalarifin12/golang_article/dto"
	"github.com/zaenalarifin12/golang_article/entity"
	"github.com/zaenalarifin12/golang_article/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	Register(dto dto.AuthRegisterDTO) entity.User
	FindByEmail(loginDTO dto.AuthLoginDTO) entity.User
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) *authService {
	return &authService{
		userRepository: userRep,
	}
}

func (services *authService) VerifyCredential(email string, password string) interface{} {
	res := services.userRepository.VerifyCredential(email, password)

	if u, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(u.Password, []byte(password))
		if u.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (services *authService) Register(registerDTO dto.AuthRegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&registerDTO))
	if err != nil {
		log.Fatalf("Failed to map")
	}
	res := services.userRepository.InsertUser(userToCreate)
	return res
}

func (services *authService) FindByEmail(loginDTO dto.AuthLoginDTO) entity.User {
	user := entity.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&loginDTO))
	if err != nil {
		log.Fatalf("Failed to map")
	}
	return services.userRepository.FindByEmail(user.Email)
}

func (services *authService) IsDuplicateEmail(email string) bool {
	res := services.userRepository.IsDuplicateEmail(email)
	if res.Error != nil {
		return true
	}
	return false
}

/**
compared password
*/

func comparePassword(password string, plainPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), plainPassword)
	if err != nil {
		return false
	}
	return true
}
