package handler

import (
	"net/http"
	"strconv"

	"backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ReplaceLineFoodHandler struct {
	Usecase usecase.ReplaceLineFoodUsecase
}

func NewReplaceLineFoodHandler(usecase usecase.ReplaceLineFoodUsecase) *ReplaceLineFoodHandler {
	return &ReplaceLineFoodHandler{Usecase: usecase}
}

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
