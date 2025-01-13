package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RestaurantHandler handles HTTP requests for restaurants.
type RestaurantHandler struct {
	RestaurantUsecase usecase.RestaurantUsecase
}

// NewRestaurantHandler creates a new instance of RestaurantHandler.
func NewRestaurantHandler(restaurantUsecase usecase.RestaurantUsecase) *RestaurantHandler {
	return &RestaurantHandler{RestaurantUsecase: restaurantUsecase}
}

// GetRestaurants handles the GET /restaurants endpoint.
func (h *RestaurantHandler) GetRestaurants(c *gin.Context) {
	restaurants, err := h.RestaurantUsecase.GetAllRestaurants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurants"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"restaurants": restaurants})
}
