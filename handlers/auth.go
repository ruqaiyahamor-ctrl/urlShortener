/* Handler - receives http requests and sends responses.

w = response writer, sends data back
r = request, contains data from the user/Postman

1. Creates auth handler
2. Registers user
3. Logs user in
4. Logs user out
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"urlShortener/models"
	"urlShortener/service"
)
//STEP 1: Store the auth service
type AuthHandler struct {
	AuthService *service.AuthService

}
//STEP 2: Create the auth handler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

//STEP 3: Register user
func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	//Decode JSON from Postman
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}
	//Call service layer
	err = h.AuthService.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}


//STEP 4:Login user
func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginRequest

	//Decode request body
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	//Call service layer
	token, err := h.AuthService.LoginUser(loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//Send token response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
//STEP 5: Logout user
func (h *AuthHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Logged out successfully. Please remove the token from Postman")
}
