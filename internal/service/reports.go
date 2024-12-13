package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type Reports struct {
	Id          int       `json:"id"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}

func (s *Service2) GetReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	reports, err := repo.RGetReportById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: reports})
}

func (s *Service2) CreateReports(c echo.Context) error {
	var reportSlice []Reports
	err := c.Bind(&reportSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	for _, report := range reportSlice {
		CreatedAT := time.Now()
		_, err = repo.CreateNewReports(report.Title, report.Description, CreatedAT.Format("2006-01-02 15:04:05"))
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "Created")
}

func (s *Service2) UpdateReportById(c echo.Context) error {
	id := c.Param("id")

	var updatedReport Reports
	err := c.Bind(&updatedReport)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.reportsRepo
	UpdatedAT := time.Now()
	_, err = repo.NewUpdateReportById(id, updatedReport.Title, updatedReport.Description, UpdatedAT.Format("2006-01-02 15:04:05"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "Updated")
}

func (s *Service2) DeleteReportById(c echo.Context) error {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	_, err = repo.DeleteReport(idInt)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "Deleted")
}
