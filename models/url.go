/* This file defines the URL data. 

1. Imports GORM
2. Creates the URL model for the database
3. It stores the original long URL
4. It stores the generated short code
5. It stores the user ID so I know who owns the URL
6. It creates ShortenRequest for the shorten URL request
*/ 

package models

import "gorm.io/gorm"

//STEP 1: Create the URL model (This struct becomes the urls table in PostgreSQL)
type URL struct {
	// STEP 2: Add GORM's default fields (This automatically gives the URL)
	gorm.Model
	
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code" gorm:"unique"`
	UserID      uint   `json:"user_id"`
}
// STEP 3: Create the shorten request model
type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}
