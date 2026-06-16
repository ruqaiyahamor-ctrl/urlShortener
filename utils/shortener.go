/* This file creates the short code for the URL. 

1. Imports math/rand
2. Creates a list of letters and numbers
3. Sets the short code length to 6 characters
4. Randomly picks 6 characters
5. Returns the finished short code
*/

package utils

import (
	"math/rand"

)
//STEP 1: Create a short code (this makes a random code)
func GenerateShortCode() string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" //STEP 2: Create the list of characters to use
	codeLength := 6
	shortCode := ""

	//STEP 2: Loop 6 times to pick a random character
	for i := 0; i < codeLength; i++ {
		randomIndex := rand.Intn(len(characters))
		shortCode += string(characters[randomIndex])
	}
	return shortCode
}