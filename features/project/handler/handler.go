package handler

import (
	"myTaskApp/app/middlewares"
	"myTaskApp/features/project"
	"myTaskApp/utils/responses"
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
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}

	requestCore := project.Core{
		UserID:      uint(idToken),
		ProjectName: newRequest.ProjectName,
		Description: newRequest.Description,
	}

	errCreate := h.projectService.Create(requestCore)
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error create project: "+errCreate.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.WebJSONResponse("success create project", nil))
}

func (h *ProjectHandler) GetAllProject(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	result, errGetAll := h.projectService.GetAll(uint(idToken))
	if errGetAll != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get all data "+errGetAll.Error(), nil))
	}

	var allProjectResponse []ProjectResponse
	for _, v := range result {
		allProjectResponse = append(allProjectResponse, ProjectResponse{
			ID:          v.ID,
			ProjectName: v.ProjectName,
		})
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get all project", allProjectResponse))
}

func (h *ProjectHandler) GetProjectById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id:"+errConv.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	result, err := h.projectService.GetProjectById(uint(idConv), uint(idToken))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error get project "+err.Error(), nil))
	}

	responseResult := ProjectResponseById{
		ID:          result.ID,
		ProjectName: result.ProjectName,
		Description: result.Description,
		TaskList:    result.TaskList,
	}
	return c.JSON(http.StatusOK, responses.WebJSONResponse("success get project", responseResult))
}

func (h *ProjectHandler) UpdateProject(c echo.Context) error {
	id := c.Param("id")
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id:"+errConv.Error(), nil))
	}

	updateRequest := UpdateRequest{}
	errBind := c.Bind(&updateRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebJSONResponse("error bind data: "+errBind.Error(), nil))
	}
	updateCore := project.Core{
		ProjectName: updateRequest.ProjectName,
		Description: updateRequest.Description,
	}

	idToken := middlewares.ExtractTokenUserId(c)
	err := h.projectService.Update(uint(idInt), uint(idToken), updateCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error update project "+err.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success update project", nil))
}

func (h *ProjectHandler) DeleteProject(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error convert id:"+errConv.Error(), nil))
	}

	idToken := middlewares.ExtractTokenUserId(c)
	tx := h.projectService.Delete(uint(idConv), uint(idToken))
	if tx != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebJSONResponse("error delete project "+tx.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebJSONResponse("success delete project", nil))
}
