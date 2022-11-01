package controllers

import (
	"net/http"
	"time"

	db "homeapp/src/db/sqlc"

	"github.com/gin-gonic/gin"
)

type FoodController struct {
	db *db.Queries
}

func NewFoodController(db *db.Queries) *FoodController {
	return &FoodController{db}
}

func (fc *FoodController) AddFood(ctx *gin.Context) {
	var food *db.Food

	if err := ctx.ShouldBindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	args := &db.AddFoodParams{
		Name:            food.Name,
		Description:     food.Description,
		Price:           food.Price,
		Calories:        "0",
		Weight:          "1",
		Amount:          "1",
		WeightPerAmount: "1",
		Photo:           "default.jpeg",
		UpdatedAt:       time.Now(),
	}

	addedFood, err := fc.db.AddFood(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"food": addedFood}})
}

func (fc *FoodController) GetFoodAvailable(ctx *gin.Context) {
	foods, err := fc.db.GetFoodAvailable(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"foods": foods}})
}
