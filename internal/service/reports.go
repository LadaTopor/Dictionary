package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetReportById ищем репорт по id
// localhost:8000/api/report/:id
func (s *Service) GetReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	report, err := repo.RGetReportById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: report})
}

// CreateReport добавляем в базу новый репорт
// localhost:8000/api/report
func (s *Service) CreateReport(c echo.Context) error {
	var report Report
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

// EditReport изменяем репорт в базе по id или добавляем новый если такого id нет
// localhost:8000/api/report/:id
func (s *Service) EditReport(c echo.Context) error {
	var report Report
	err := c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.reportsRepo
	err = repo.EditReport(report.Title, report.Description, id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// DeleteReport удаляем репорт из базы по id
// localhost:8000/api/report/:id
func (s *Service) DeleteReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
	}
	repo := s.reportsRepo
	err = repo.DeleteReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.String(http.StatusOK, "OK")
}
