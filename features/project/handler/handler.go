package handler

import (
	"myTaskApp/features/project"
	"net/http"

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
	newRequest := ProjectRequest{}
	errBind := c.Bind(&newRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	requestCore := project.Core{
		UserID:      newRequest.UserID,
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
	requestAll := ProjectResponse{}
	errBind := c.Bind(&requestAll)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	result, errGetAll := h.projectService.GetAll()
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
			UserID:      v.UserID,
			ProjectName: v.ProjectName,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success get all project",
		"project": allProjectResponse,
	})
}
