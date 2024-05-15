package handler

import (
	"myTaskApp/app/middlewares"
	"myTaskApp/features/task"
	"myTaskApp/utils/responses"
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
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
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
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create task: "+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success create task", nil))
}

func (h *TaskHandler) UpdateTaskById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id: "+errConv.Error(), nil))
	}

	updateRequest := TaskUpdateRequest{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	updateCore := task.Core{
		TaskName:        updateRequest.TaskName,
		DescriptionTask: updateRequest.DescriptionTask,
		StatusTask:      updateRequest.StatusTask,
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := h.HandlerService.Update(uint(idConv), uint(idToken), updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update task: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success update task by id", nil))
}

func (h *TaskHandler) DeleteTaskById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id: "+errConv.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := h.HandlerService.Delete(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error delete task: "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success delete task", nil))
}
