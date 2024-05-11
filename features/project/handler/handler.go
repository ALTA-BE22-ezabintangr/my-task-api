package handler

import (
	"myTaskApp/app/middlewares"
	"myTaskApp/features/project"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectService project.ServiceInterface
}

func New(ph project.ServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		projectService: ph,
	}
}

func (h *ProjectHandler) CreateProject(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	newRequest := ProjectRequest{}
	errBind := c.Bind(&newRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	requestCore := project.Core{
		UserID:      uint(idToken),
		ProjectName: newRequest.ProjectName,
		Description: newRequest.Description,
	}

	errCreate := h.projectService.Create(requestCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error create data " + errCreate.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success create project",
	})
}

func (h *ProjectHandler) GetAllProject(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, errGetAll := h.projectService.GetAll(uint(idToken))
	if errGetAll != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error get all data " + errGetAll.Error(),
		})
	}

	var allProjectResponse []ProjectResponse
	for _, v := range result {
		allProjectResponse = append(allProjectResponse, ProjectResponse{
			ID:          v.ID,
			ProjectName: v.ProjectName,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success get all project",
		"project": allProjectResponse,
	})
}

func (h *ProjectHandler) GetProjectById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error convert id " + errConv.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	result, err := h.projectService.GetProjectById(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error get data " + err.Error(),
		})
	}

	responseResult := ProjectResponseById{
		ID:          result.ID,
		ProjectName: result.ProjectName,
		Description: result.Description,
		TaskList:    result.TaskList,
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success get project",
		"project": responseResult,
	})
}

func (h *ProjectHandler) UpdateProject(c echo.Context) error {
	id := c.Param("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error convert id " + errConv.Error(),
		})
	}

	updateRequest := UpdateRequest{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data " + errBind.Error(),
		})
	}
	updateCore := project.Core{
		ProjectName: updateRequest.ProjectName,
		Description: updateRequest.Description,
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := h.projectService.Update(uint(idInt), uint(idToken), updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error update data" + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "succes update project",
	})
}

func (h *ProjectHandler) DeleteProject(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error convert id " + errConv.Error(),
		})
	}

	idToken := middlewares.ExtractTokenUserId(c)
	tx := h.projectService.Delete(uint(idConv), uint(idToken))
	if tx != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete data " + tx.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success delete project",
	})
}
