package routes

import (
	"myTaskApp/features/user/data"
	"myTaskApp/features/user/handler"
	"myTaskApp/features/user/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	dataService := data.New(db)
	userService := service.New(dataService)
	userHandlerAPI := handler.New(userService)

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "hello world",
		})
	})

	e.POST("/users", userHandlerAPI.Register)
}
