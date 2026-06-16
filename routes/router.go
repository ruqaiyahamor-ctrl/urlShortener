/* Routes - connects endpoints to handlers.

1. Creates repositories
2. Creates services
3. Creates handlers
4. Creates router
5. Adds public routes
6. Adds protected routes
7. Returns router
*/

package routes

import (
	"urlShortener/handlers"
	"urlShortener/middleware"
	"urlShortener/repository"
	"urlShortener/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *mux.Router {
	//STEP 1: Create repositories
	userRepo := repository.NewUserRepository(db)
	urlRepo := repository.NewURLRepository(db)

	//STEP 2: Create services
	authService := service.NewAuthService(userRepo)
	urlService := service.NewURLService(urlRepo)

	//STEP 3: Create handlers
	authHandler := handlers.NewAuthHandler(authService)
	urlHandler := handlers.NewURLHandler(urlService)

	//STEP 4: Create router
	r := mux.NewRouter()

	//Step 5: Public routes
	r.HandleFunc("/register", authHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", authHandler.LoginUser).Methods("POST")
	r.HandleFunc("/logout", authHandler.LogoutUser).Methods("POST")
	r.HandleFunc("/{code}", urlHandler.RedirectURL).Methods("GET")

	//STEP 6: Protected routes
	protectedRoutes := r.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	protectedRoutes.HandleFunc("/shorten", urlHandler.ShortenURL).Methods("POST")
	protectedRoutes.HandleFunc("/url/{code}", urlHandler.DeleteURL).Methods("DELETE")

	return r
}
