/* Handler - receives http requests and sends responses.

1. Creates URL handler
2. Shortens URL
3. Redirects short URL
4. Deletes URL
*/

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"urlShortener/middleware"
	"urlShortener/models"
	"urlShortener/service"

	"github.com/gorilla/mux"
)

// STEP 1: Store the URL service
type URLHandler struct {
	URLService *service.URLService
}

// STEP 2: Create the URL handler
func NewURLHandler(urlService *service.URLService) *URLHandler {
	return &URLHandler{URLService: urlService}
}

// STEP 3: Shorten the URL
func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request models.ShortenRequest

	//Decode request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//Get logged-in user ID
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Call service layer
	shortURL, err := h.URLService.ShortenURL(request, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Return short URL
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"short_url": shortURL,
	})
}

// STEP 4: Redirect short URL
func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	//Get {code} from the URL path
	vars := mux.Vars(r)
	shortCode := vars["code"]

	//Find original URL
	originalURL, err := h.URLService.GetOriginalURL(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	//Redirect to the original URL
	http.Redirect(w, r, originalURL, http.StatusFound)
}

// STEP 5: Delete URL
func (h *URLHandler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["code"]

	//Get logged in user ID
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//Call service layer
	err = h.URLService.DeleteURL(shortCode, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Send cussess response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "URL deleted successfully")

}
