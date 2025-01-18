package handler

import (
	"net/http"
	"strconv"

	"backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

// CreateLineFoodHandler handles the POST /line_foods/:food_id request.
type CreateLineFoodHandler struct {
	Usecase usecase.CreateLineFoodUsecase
}

// NewCreateLineFoodHandler creates a new instance of CreateLineFoodHandler.
func NewCreateLineFoodHandler(usecase usecase.CreateLineFoodUsecase) *CreateLineFoodHandler {
	return &CreateLineFoodHandler{Usecase: usecase}
}

// Handle processes the request to create or update a line food.
func (h *CreateLineFoodHandler) Handle(c *gin.Context) {
	foodID, err := strconv.Atoi(c.Param("food_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food_id"})
		return
	}

	count, err := strconv.Atoi(c.PostForm("count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid count"})
		return
	}

	lineFood, conflictInfo, err := h.Usecase.Execute(foodID, count)
	if err != nil {
		if conflictInfo != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"existing_restaurant": conflictInfo["existing_restaurant"],
				"new_restaurant":      conflictInfo["new_restaurant"],
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"line_food": lineFood})
}
