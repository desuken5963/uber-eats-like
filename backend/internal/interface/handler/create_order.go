package handler

import (
	"backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderHandler struct {
	usecase usecase.CreateOrderUsecase
}

func NewCreateOrderHandler(u usecase.CreateOrderUsecase) *CreateOrderHandler {
	return &CreateOrderHandler{usecase: u}
}

func (h *CreateOrderHandler) Handle(c *gin.Context) {
	var req struct {
		LineFoodIDs []int `json:"line_food_ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.usecase.Execute(req.LineFoodIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
