/*Service - main business logic.
 
1. Creates URL service
2. Validates the original URL
3. Generates a short code
4. Saves the URL to the database
5. Returns the short URL
6. Finds original URL for redirect
7. Deletes a URL
*/

package service

import (
	"fmt"
	"os"
	"strings"
	"urlShortener/models"
	"urlShortener/repository"
	"urlShortener/utils"
)
//STEP 1: Store the URL repository
type URLService struct {
	URLRepo repository.URLRepository
}
//STEP 2: Create the URL service
func NewURLService(urlRepo repository.URLRepository) *URLService {
	return &URLService{URLRepo: urlRepo}
}

//STEP 3: Shorten a URL
func (s *URLService) ShortenURL(request models.ShortenRequest, userID uint) (string, error) {
	//Validate original URL
	if request.OriginalURL == "" {
		return "", fmt.Errorf("Original URL is request")
	}
	//Check the URL starts correctly
	if !strings.HasPrefix(request.OriginalURL, "http://") && !strings.HasPrefix(request.OriginalURL, "https://") {
		return "", fmt.Errorf("URL should start with http:// or https://")
	}

	//Generate short code
	shortCode := utils.GenerateShortCode()

	//Create URL record
	url := models.URL{
		OriginalURL: request.OriginalURL,
		ShortCode:   shortCode,
		UserID:      userID,
	}
	//Save url into database
	err := s.URLRepo.CreateURL(url)
	if err != nil {
		return "", err
	}

	//Build full short URL
	baseURL := os.Getenv("BASE_URL")
	shortURL := baseURL + "/" + shortCode

	return shortURL, nil
}

//STEP 4: Find original URL for redirect
func (s *URLService) GetOriginalURL(ShortCode string) (string, error) {
	urlObj, err := s.URLRepo.GetURLByCode(ShortCode)
	if err != nil {
		return "", fmt.Errorf("URL not found")
	}

	return urlObj.OriginalURL, nil
}

//STEP 5: Delete a URL
func (s *URLService) DeleteURL(shortCode string, userID uint) error {

	//Find URL by short code
	url, err := s.URLRepo.GetURLByCode(shortCode)
	if err != nil {
		return fmt.Errorf("URL not found")
	}

	//Check the URL belongs to the logged-in user
	if url.UserID != userID {
		return fmt.Errorf("you cannot delete this URL")
	}

	//Delete URL
	return s.URLRepo.DeleteURL(url)
}
