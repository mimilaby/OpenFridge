// Package controllers handle the request and response
package controllers

import (
	"net/http"
	"time"

	db "openfridge/src/db/sqlc"
	"openfridge/src/utils"

	"github.com/gin-gonic/gin"
)

// AuthController holds connection to the database
type AuthController struct {
	db *db.Queries
}

// NewAuthController connect Controller to the database
func NewAuthController(db *db.Queries) *AuthController {
	return &AuthController{db}
}

// SignUpUser adds a new user to the database, with Name, Email and Password parameters
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var credentials *db.User

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword := utils.HashPassword(credentials.Password)

	args := &db.CreateUserParams{
		Name:      credentials.Name,
		Email:     credentials.Email,
		Password:  hashedPassword,
		Photo:     "default.jpeg",
		Verified:  true,
		Role:      "user",
		UpdatedAt: time.Now(),
	}

	user, err := ac.db.CreateUser(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": user}})
}
