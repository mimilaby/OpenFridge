// Package routes handles all the routes for the application
package routes

import (
	"openfridge/src/controllers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes struct holds the auth controller
type AuthRoutes struct {
	authController controllers.AuthController
}

// NewAuthRoutes connects the auth controller to the auth routes
func NewAuthRoutes(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

// AuthRoute add the auth routes to the router
func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {

	router := rg.Group("/auth")
	router.POST("/register", rc.authController.SignUpUser)
}
