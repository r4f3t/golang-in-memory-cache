package controller

import (
	"github.com/r4f3t/webapi/internal/singletonObjectCache/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type resource struct {
	service user.Service
}

func NewController(service user.Service) *resource {
	return &resource{
		service: service,
	}
}

func (receiver *resource) getUser(c echo.Context) error {
	idString := c.QueryParam("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Unsupported Request")
	}

	result := receiver.service.GetUserById(id)

	if result == nil {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}

	return c.JSON(http.StatusOK, result)
}
