package routes

import (
	_projectData "myTaskApp/features/project/data"
	_projectHandler "myTaskApp/features/project/handler"
	_projectService "myTaskApp/features/project/service"
	"myTaskApp/features/task/data"
	"myTaskApp/features/task/handler"
	"myTaskApp/features/task/service"
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

	projectData := _projectData.New(db)
	projectService := _projectService.New(projectData)
	projectHandlerAPI := _projectHandler.New(projectService)

	taskData := data.New(db)
	taskService := service.New(taskData)
	taskHandlerAPI := handler.New(taskService)

	e.POST("/users", userHandlerAPI.Register)
	e.GET("/users", userHandlerAPI.GetAll)
	e.DELETE("/users/:id", userHandlerAPI.Delete)
	e.PUT("/users/:id", userHandlerAPI.Update)

	e.POST("/login", userHandlerAPI.Login)

	e.POST("/projects", projectHandlerAPI.CreateProject)
	e.GET("/projects", projectHandlerAPI.GetAllProject)
	e.PUT("/projects/:id", projectHandlerAPI.UpdateProject)
	e.DELETE("/projects/:id", projectHandlerAPI.DeleteProject)

	e.POST("/tasks", taskHandlerAPI.CreateTask)
}
