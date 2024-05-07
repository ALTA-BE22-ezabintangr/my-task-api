package routes

import (
	"myTaskApp/features/user/data"
	"myTaskApp/features/user/handler"
	"myTaskApp/features/user/service"
	encrypts "myTaskApp/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	hashService := encrypts.NewHashService()
	dataService := data.New(db)
	userService := service.New(dataService, hashService)
	userHandlerAPI := handler.New(userService)

	e.POST("/users", userHandlerAPI.Register)
	e.GET("/users", userHandlerAPI.GetAll)
	e.DELETE("/users/:id", userHandlerAPI.Delete)
	e.PUT("/users/:id", userHandlerAPI.Update)

	e.POST("/login", userHandlerAPI.Login)

}
