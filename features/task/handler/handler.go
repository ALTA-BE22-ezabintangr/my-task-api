package handler

import (
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

func (h *TaskHandler) GetTaskbyUserId(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}

	result, err := h.HandlerService.GetTaskbyId(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error get task " + err.Error(),
		})
	}

	var resultResponse []ResponseById
	for _, v := range result {
		resultResponse = append(resultResponse, ResponseById{
			ProjectID:  v.ProjectID,
			TaskName:   v.TaskName,
			StatusTask: v.StatusTask,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success get task by id user",
		"result":  resultResponse,
	})
}
