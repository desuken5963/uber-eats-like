package handler

import (
	"net/http"
	"strconv"

	"backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

// ReplaceLineFoodHandler handles the PUT /line_foods/:food_id request.
type ReplaceLineFoodHandler struct {
	Usecase usecase.ReplaceLineFoodUsecase
}

// NewReplaceLineFoodHandler creates a new instance of ReplaceLineFoodHandler.
func NewReplaceLineFoodHandler(usecase usecase.ReplaceLineFoodUsecase) *ReplaceLineFoodHandler {
	return &ReplaceLineFoodHandler{Usecase: usecase}
}

// Handle processes the request to replace line foods.
func (h *ReplaceLineFoodHandler) Handle(c *gin.Context) {
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

	lineFood, err := h.Usecase.Execute(foodID, count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"line_food": lineFood})
}
