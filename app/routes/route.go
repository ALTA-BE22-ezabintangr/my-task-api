package routes

import (
	"myTaskApp/app/middlewares"
	_projectData "myTaskApp/features/project/data"
	_projectHandler "myTaskApp/features/project/handler"
	_projectService "myTaskApp/features/project/service"
	_taskData "myTaskApp/features/task/data"
	_taskHandler "myTaskApp/features/task/handler"
	_taskService "myTaskApp/features/task/service"
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

	taskData := _taskData.New(db)
	taskService := _taskService.New(taskData, projectData)
	taskHandlerAPI := _taskHandler.New(taskService)

	e.POST("/login", userHandlerAPI.Login)

	e.POST("/users", userHandlerAPI.Register)
	e.GET("/users", userHandlerAPI.GetProfileUser, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.PUT("/users", userHandlerAPI.Update, middlewares.JWTMiddleware())

	e.POST("/projects", projectHandlerAPI.CreateProject, middlewares.JWTMiddleware())
	e.GET("/projects", projectHandlerAPI.GetAllProject, middlewares.JWTMiddleware())
	e.GET("/projects/:id", projectHandlerAPI.GetProjectById, middlewares.JWTMiddleware())
	e.PUT("/projects/:id", projectHandlerAPI.UpdateProject, middlewares.JWTMiddleware())
	e.DELETE("/projects/:id", projectHandlerAPI.DeleteProject, middlewares.JWTMiddleware())

	e.POST("/tasks", taskHandlerAPI.CreateTask, middlewares.JWTMiddleware())
	e.PUT("/tasks/:id", taskHandlerAPI.UpdateTaskById, middlewares.JWTMiddleware())
	e.DELETE("/tasks/:id", taskHandlerAPI.DeleteTaskById, middlewares.JWTMiddleware())
}
