package handler

import (
	"myTaskApp/app/middlewares"
	"myTaskApp/features/task"
	"net/http"
	"strconv"

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
	idToken := middlewares.ExtractTokenUserId(c)
	newTaskCore := task.Core{
		UserID:          uint(idToken),
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

func (h *TaskHandler) UpdateTaskById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id " + errConv.Error(),
		})
	}

	updateRequest := TaskUpdateRequest{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data " + errBind.Error(),
		})
	}

	updateCore := task.Core{
		TaskName:        updateRequest.TaskName,
		DescriptionTask: updateRequest.DescriptionTask,
		StatusTask:      updateRequest.StatusTask,
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := h.HandlerService.Update(uint(idConv), uint(idToken), updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error update data" + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success update task by id",
	})
}

func (h *TaskHandler) DeleteTaskById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id " + errConv.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := h.HandlerService.Delete(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete data" + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success delete task by id",
	})
}
