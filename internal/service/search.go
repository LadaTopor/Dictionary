package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s *Service) SearchByTitle(c echo.Context) error {
	title := c.QueryParam("title")
	repo := s.searchRepo
	result, err := repo.SearchTitle(title)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, result)
}
