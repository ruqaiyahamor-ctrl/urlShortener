/* Service - main business logic.

1. Creates auth service
2. Checks if user exists
3. Registers user
4. Logins user
*/

package service

import (
	"urlShortener/middleware"
	"fmt"
	"urlShortener/models"
	"urlShortener/repository"
	"urlShortener/utils"
)
//STEP 1: Store the user repository
type AuthService struct {
	UserRepo repository.UserRepository
}
//STEP 2: Create the auth service
func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

//STEP 3: Check if email already exists
func (s *AuthService) CheckUserExists(email string) error {
	err := s.UserRepo.CheckUserExists(email)
	if err == nil {
		return fmt.Errorf("user already exists with email %s", email)
	}
	return nil
}
//STEP 4: Register a new user
func (s *AuthService) RegisterUser(user models.User) error {
	err := s.CheckUserExists(user.Email)
	if err != nil {
		return err
	}
	//Hash password
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	//Add hashed password into the user
	user.Password = hashPassword

	//Save user into database
	return s.UserRepo.CreateUser(user)
}

//STEP 5: Login user
func (s *AuthService) LoginUser(loginRequest models.LoginRequest) (string, error) {
	//Get user by email
	user, err := s.UserRepo.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	// Check password
	err = utils.ComparePasswords(loginRequest.Password, user.Password)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	//Generate token
	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err

	}
	return token, nil
}