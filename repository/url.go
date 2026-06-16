/* Repository talks to the database.

1. Defines URL database actions
2. Stores the database connection
3. Creates the repository
4. Saves a shortened URL
5. Finds URL by short code
6. Deletes a URL
*/

package repository

import (
	
	"urlShortener/models"

	"gorm.io/gorm"
)


//STEP 1: List the URL database actions
type URLRepository interface {
	CreateURL(url models.URL) error
	GetURLByCode (shortCode string)(models.URL, error)
	DeleteURL(url models.URL) error
}
//STEP 2: Store the database connection
type PostgresURLRepo struct {
	db *gorm.DB
}
//STEP 3: Create the URL repository
func NewURLRepository(db *gorm.DB) URLRepository {
	return &PostgresURLRepo{db: db}

}

//STEP 4: Save shortened URL
func (r PostgresURLRepo) CreateURL(url models.URL) error {
	err := r.db.Create(&url).Error
	return err
}

//STEP 5: Find URL by it's short code
func (r PostgresURLRepo) GetURLByCode(shortCode string) (models.URL, error) {
	var url models.URL

	err := r.db.Where("short_code =?", shortCode).First(&url).Error
	if err != nil {
		return models.URL{}, err
	}
	return url, nil
}
//STEP 6: Delete URL from the databse
func (r PostgresURLRepo) DeleteURL(url models.URL) error {
	err := r.db.Delete(&url).Error
	return err
}