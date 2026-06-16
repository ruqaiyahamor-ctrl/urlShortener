/* This file handles password security. 

1. Imports bcrypt
2. Hashes a plain password before saving it to the database
3. Then compares the login password with the hashed password in the database
*/

package utils

import "golang.org/x/crypto/bcrypt"

//STEP 1: Hash the user's password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

//STEP 2: Compare the login password with the saved hashed password
func ComparePasswords(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}