/* This defines the user data. 

1. Imports GORM
2. Creates the User model for the database
3. Adds JSON names for Postman requests
4. Makes the email unique
5. Creates LoginRequest for login only
*/

package models

import "gorm.io/gorm"

//STEP 1: Create the user model (This struct becomes the users table in PostgresSQL)
type User struct {
	gorm.Model //STEP 2: Add GORMS's default field
	//STEP 3: Store the user's details in Postman
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
// STEP 4: Create the login request model (used when user log's in)
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}