package routes

import (
	"myTaskApp/features/user/data"
	"myTaskApp/features/user/handler"
	"myTaskApp/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	dataService := data.New(db)
	userService := service.New(dataService)
	userHandlerAPI := handler.New(userService)

	e.GET("/users", userHandlerAPI.GetAll)
	e.POST("/users", userHandlerAPI.Register)
}
