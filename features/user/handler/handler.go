package handler

import (
	"myTaskApp/features/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.ServiceInterface
}

func New(us user.ServiceInterface) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) Register(c echo.Context) error {
	newUser := UserRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	inputCore := user.Core{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		Phone:    newUser.Phone,
		Address:  newUser.Address,
	}
	errInsert := uh.userService.Create(inputCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error insert data " + errInsert.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success add user",
	})
}

func (uh *UserHandler) GetAll(c echo.Context) error {
	result, err := uh.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error read data " + err.Error(),
		})
	}

	var allUserResponse []UserResponse
	for _, value := range result {
		allUserResponse = append(allUserResponse, UserResponse{
			ID:    value.ID,
			Name:  value.Name,
			Email: value.Email,
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success read data",
		"results": allUserResponse,
	})
}

func (uh *UserHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + err.Error(),
		})
	}

	tx := uh.userService.Delete(uint(idConv))
	if tx != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error delete data " + tx.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "success delete user",
	})
}

func (uh *UserHandler) Update(c echo.Context) error {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + err.Error(),
		})
	}

	updateUser := UserRequest{}
	errBind := c.Bind(&updateUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	updateCore := user.Core{
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Password: updateUser.Password,
		Phone:    updateUser.Phone,
		Address:  updateUser.Address,
	}

	errUpdate := uh.userService.Update(uint(idConv), updateCore)
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error update data " + errUpdate.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success update user",
	})
}

func (uh *UserHandler) Login(c echo.Context) error {
	loginUser := LoginRequest{}
	errBind := c.Bind(&loginUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error bind data: " + errBind.Error(),
		})
	}

	login, errLogin := uh.userService.Login(loginUser.Email, loginUser.Password)
	if errLogin != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": "error login " + errLogin.Error(),
		})
	}

	var resultResponse = map[string]any{
		"id":    login.ID,
		"name":  login.Name,
		"email": login.Email,
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success login",
		"data":    resultResponse,
	})
}
