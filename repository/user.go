/* Repository - talks to the database.

1. Defines user database actions
2. Stores the database connection
3. Creates the repository
4. Checks if email exists
5. Saves a new user
6. Finds user by email
*/

package repository

import (
	"urlShortener/models"

	"gorm.io/gorm"
)
//STEP 1: List of user database actions
type UserRepository interface {
	CheckUserExists(email string) error
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
}
//STEP 2: Store the database connection
type PostgresUserRepo struct {
	db *gorm.DB

}
//STEP 3: Create the user repository
func NewUserRepository (db *gorm.DB) UserRepository {
	return &PostgresUserRepo{db: db}
}

//STEP 4: Check if user email already exists
func (r PostgresUserRepo) CheckUserExists(email string) error {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	return err
}

//STEP 5: Save the new user
func (r PostgresUserRepo) CreateUser(user models.User) error {

	err := r.db.Create(&user).Error
	return err
}

//STEP 6: Find user by email for login
func (r PostgresUserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
