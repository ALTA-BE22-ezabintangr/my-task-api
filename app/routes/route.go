package routes

import (
	"myTaskApp/features/project/data"
	"myTaskApp/features/project/handler"
	"myTaskApp/features/project/service"
	_userData "myTaskApp/features/user/data"
	_userHandler "myTaskApp/features/user/handler"
	_userService "myTaskApp/features/user/service"
	encrypts "myTaskApp/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {

	hashService := encrypts.NewHashService()
	dataService := _userData.New(db)
	userService := _userService.New(dataService, hashService)
	userHandlerAPI := _userHandler.New(userService)

	projectData := data.New(db)
	projectService := service.New(projectData)
	projectHandlerAPI := handler.New(projectService)

	e.POST("/users", userHandlerAPI.Register)
	e.GET("/users", userHandlerAPI.GetAll)
	e.DELETE("/users/:id", userHandlerAPI.Delete)
	e.PUT("/users/:id", userHandlerAPI.Update)

	e.POST("/login", userHandlerAPI.Login)

	e.POST("/projects", projectHandlerAPI.CreateProject)
	e.GET("/projects", projectHandlerAPI.GetAllProject)
	e.PUT("/projects/:id", projectHandlerAPI.UpdateProject)
	e.DELETE("/projects/:id", projectHandlerAPI.DeleteProject)
}
