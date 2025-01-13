package handler

import (
	"net/http"
	"strconv"

	"backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

// FoodHandler handles HTTP requests for foods.
type FoodHandler struct {
	FoodUsecase usecase.FoodUsecase
}

// NewFoodHandler creates a new instance of FoodHandler.
func NewFoodHandler(foodUsecase usecase.FoodUsecase) *FoodHandler {
	return &FoodHandler{FoodUsecase: foodUsecase}
}

// GetFoods handles the GET /restaurants/:restaurant_id/foods endpoint.
func (h *FoodHandler) GetFoods(c *gin.Context) {
	restaurantID, err := strconv.Atoi(c.Param("restaurant_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant_id"})
		return
	}

	foods, err := h.FoodUsecase.GetFoodsByRestaurantID(restaurantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch foods"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"foods": foods})
}
