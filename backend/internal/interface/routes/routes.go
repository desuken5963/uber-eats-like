package routes

import (
	"backend/internal/config"
	"backend/internal/interface/handler"
	"backend/internal/interface/repository"
	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the routes and their dependencies.
func RegisterRoutes(router *gin.Engine) {
	// データベース接続
	db := config.DB

	// 依存性注入
	restaurantRepo := repository.NewRestaurantRepository(db)
	restaurantUsecase := usecase.NewRestaurantUsecase(restaurantRepo)
	restaurantHandler := handler.NewRestaurantHandler(restaurantUsecase)

	foodRepo := repository.NewFoodRepository(db)
	foodUsecase := usecase.NewFoodUsecase(foodRepo)
	foodHandler := handler.NewFoodHandler(foodUsecase)

	getlineFoodsRepo := repository.NewGetLineFoodsRepository(db)
	getLineFoodsUsecase := usecase.NewGetLineFoodsUsecase(getlineFoodsRepo)
	getLineFoodsHandler := handler.NewGetLineFoodsHandler(getLineFoodsUsecase)

	createLineFoodRepo := repository.NewCreateLineFoodRepository(db)
	createLineFoodUsecase := usecase.NewCreateLineFoodUsecase(createLineFoodRepo)
	createLineFoodHandler := handler.NewCreateLineFoodHandler(createLineFoodUsecase)

	replaceLineFoodRepo := repository.NewReplaceLineFoodRepository(db)
	replaceLineFoodUsecase := usecase.NewReplaceLineFoodUsecase(replaceLineFoodRepo)
	replaceLineFoodHandler := handler.NewReplaceLineFoodHandler(replaceLineFoodUsecase)

	// ルート設定
	api := router.Group("/api/v1")
	{
		api.GET("/restaurants", restaurantHandler.GetRestaurants)
		api.GET("/restaurants/:restaurant_id/foods", foodHandler.GetFoods)
		api.GET("/line_foods", getLineFoodsHandler.Handle)
		api.POST("/line_foods/:food_id", createLineFoodHandler.Handle)
		api.PUT("/line_foods/:food_id", replaceLineFoodHandler.Handle)
	}
}
