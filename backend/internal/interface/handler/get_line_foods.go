package handler

import (
	"net/http"

	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

// GetLineFoodsHandler handles the GET /line_foods request.
type GetLineFoodsHandler struct {
	Usecase usecase.GetLineFoodsUsecase
}

// NewGetLineFoodsHandler creates a new instance of GetLineFoodsHandler.
func NewGetLineFoodsHandler(usecase usecase.GetLineFoodsUsecase) *GetLineFoodsHandler {
	return &GetLineFoodsHandler{Usecase: usecase}
}

// Handle processes the request and returns active line foods.
func (h *GetLineFoodsHandler) Handle(c *gin.Context) {
	response, err := h.Usecase.Execute()
	if err != nil {
		if err.Error() == "no content" {
			c.JSON(http.StatusNoContent, gin.H{})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch line foods"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}
