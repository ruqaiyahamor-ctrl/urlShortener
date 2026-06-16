/* Middleware checks login/auth. This file handles JWT authentication.

1. Checks if protected routes have a token
2. Blocks requests without a valid token
3. Creates a JWT token after login
4. Checks if a JWT token is real/valid
5. Gets the logged-in user's ID from the token
*/

package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

//STEP 1: AuthMiddleware runs before protected routes like api/shorten and api/url/{code}. It checks if the request has a valid Authorization token.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		//STEP 2: Get the token from the request header (postman/authorization/bearer token)
		tokenString := r.Header.Get("Authorization") 
		//STEP 3: if no token, block request
		if tokenString == "" { 
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		//STEP 4: Remove Bearer from the token string. This leaves the JWT token.
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		//STEP 5: Check if the token is valid
		_, err := VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "invalide token", http.StatusUnauthorized)
			return
		}
		//STEP 6: If the token is valid, continue to protected route
		next.ServeHTTP(w, r)
	})
}

//STEP 7: Create a JWT token after login (used in the login service after the email and password are correct)
func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims { 
		"userID": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), 
	})
	//STEP 8: Get the JWT secret from the .env file
	jwtSecret := os.Getenv("JWT_SECRET")

	//STEP 9: Sign the token and turn it into a string
	tokenString, err := token.SignedString([]byte(jwtSecret)) 
	if err != nil {
		return "", err
	}
	//STEP 10: Return the finished token
	return tokenString, nil
}


//STEP 11: VerifyJWT will check if the token is valid
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	//STEP 12: Parse/check the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//STEP 13: Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}


//STEP 14: GetUserIDFromToken will get the logged-in user's ID from the token
func GetUserIDFromToken(r *http.Request) (uint, error) {
	authHeader := r.Header.Get("Authorization")

	//STEP 15: If the header is missing, return an error
	if authHeader == "" {
		return 0, fmt.Errorf("authorization header is required")
	}
	//STEP 16: Remove Bearer so it only keeps the token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	//STEP 17: Verify the token first
	parsedToken, err := VerifyJWT(tokenString)
	if err != nil {
		return 0, fmt.Errorf("invalid token")
	}
	//STEP 18: Read the claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}
	//STEP 19: Get the userID from the token claims - JWT numbers come back as float64
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, fmt.Errorf("user ID not found in token")
	}
	//STEP 20: Convert the userID from float64 to uint (because GORMs ID is uint)
	return uint(userIDFloat), nil

}
