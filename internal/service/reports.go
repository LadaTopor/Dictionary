package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *Service) GetReportById(c echo.Context) error {
	reportIds := c.QueryParam("id")
	ids := strings.Split(reportIds, ",")
	var err error

	var idsInt []int
	if len(reportIds) > 0 {
		for i := 0; i < len(ids); i++ {
			idInt, err := strconv.Atoi(ids[i])
			if err != nil {
				s.logger.Error(err)
				return err
			}
			idsInt = append(idsInt, idInt)
		}
	}
	reports, err := s.reportsRepo.RGetReportById(idsInt)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: reports})
}

func (s *Service) CreateReport(c echo.Context) error {
	var reportSlice Reports
	err := c.Bind(&reportSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	report := reportSlice
	err = repo.CreateNewReport(report.Title, report.Description, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "Created")
}

func (s *Service) UpdateReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	var updatedReport Reports
	err = c.Bind(&updatedReport)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.reportsRepo
	_, err = repo.NewUpdateReportById(id, updatedReport.Title, updatedReport.Description, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "Updated")
}

func (s *Service) DeleteReportById(c echo.Context) error {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	err = repo.DeleteReport(idInt)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "Deleted")
}
