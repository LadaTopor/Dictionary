package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetReportById ищем слово по id
// localhost:8000/api/report/:id
func (s *Service) GetReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	report, err := repo.GetReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: report})
}

// CreateReport ищем слово по id
// localhost:8000/api/report
func (s *Service) CreateReport(c echo.Context) error {
	report := new(Report)
	err := c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	err = repo.CreateNewReport(report.Title, report.Description)

	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// UpdateReportById ищем слово по id
// localhost:8000/api/report/:id
func (s *Service) UpdateReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	report := new(Report)
	err = c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	err = repo.UpdateReport(id, report.Title, report.Description)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "OK")
}

// DeleteReportById удаляем слово по id
// localhost:8000/api/report/:id
func (s *Service) DeleteReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	err = repo.DeleteReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, "OK")
}
