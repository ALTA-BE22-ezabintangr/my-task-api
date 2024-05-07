package handler

import (
	"myTaskApp/features/task"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	HandlerService task.ServiceInterface
}

func New(th task.ServiceInterface) *TaskHandler {
	return &TaskHandler{
		HandlerService: th,
	}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	newTask := TaskRequest{}
	errBind := c.Bind(&newTask)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data" + errBind.Error(),
		})
	}

	newTaskCore := task.Core{
		UserID:          newTask.UserID,
		ProjectID:       newTask.ProjectID,
		TaskName:        newTask.TaskName,
		DescriptionTask: newTask.DescriptionTask,
		StatusTask:      newTask.StatusTask,
	}

	errCreate := h.HandlerService.Create(newTaskCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error create task" + errCreate.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success create task",
	})
}
