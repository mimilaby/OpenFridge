package routes

import (
	"homeapp/src/controllers"

	"github.com/gin-gonic/gin"
)

type FoodRoutes struct {
	foodController controllers.FoodController
}

func NewFoodRoutes(foodController controllers.FoodController) FoodRoutes {
	return FoodRoutes{foodController}
}

func (rc *FoodRoutes) FoodRoute(rg *gin.RouterGroup) {

	router := rg.Group("/food")
	router.POST("/add", rc.foodController.AddFood)
}
