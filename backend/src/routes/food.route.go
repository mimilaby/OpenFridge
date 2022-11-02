package routes

import (
	"homeapp/src/controllers"

	"github.com/gin-gonic/gin"
)

// FoodRoutes struct holds the food controller
type FoodRoutes struct {
	foodController controllers.FoodController
}

// NewFoodController connects the food controller to the food routes
func NewFoodRoutes(foodController controllers.FoodController) FoodRoutes {
	return FoodRoutes{foodController}
}

// FoodRoute add the food routes to the router
func (rc *FoodRoutes) FoodRoute(rg *gin.RouterGroup) {

	router := rg.Group("/food")
	router.POST("/add", rc.foodController.AddFood)
	router.GET("/get_available", rc.foodController.GetFoodAvailable)
}
