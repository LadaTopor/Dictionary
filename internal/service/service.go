package service

import (
	"database/sql"
	"dictionary/internal/reports"

	"dictionary/internal/words"

	"github.com/labstack/echo/v4"
)

const (
	InvalidParams       = "invalid params"
	InternalServerError = "internal error"
)

type Service struct {
	db     *sql.DB
	logger echo.Logger

	wordsRepo *words.Repo
}

type Service2 struct {
	db     *sql.DB
	logger echo.Logger

	reportsRepo *reports.Repo
}

func NewService(db *sql.DB, logger echo.Logger) *Service {
	svc := &Service{
		db:     db,
		logger: logger,
	}
	svc.initRepositories(db)

	return svc
}

func NewService2(db *sql.DB, logger echo.Logger) *Service2 {
	svc := &Service2{
		db:     db,
		logger: logger,
	}
	svc.initRepositories(db)

	return svc
}

func (s *Service) initRepositories(db *sql.DB) {
	s.wordsRepo = words.NewRepo(db)
}
func (s *Service2) initRepositories(db *sql.DB) {
	s.reportsRepo = reports.NewRepo(db)
}

// Пока можно не вдаваться в то что ниже

type Response struct {
	Object       any    `json:"object,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func (r *Response) Error() string {
	return r.ErrorMessage
}

func (s *Service) NewError(err string) (int, *Response) {
	return 400, &Response{ErrorMessage: err}
}

func (s *Service2) NewError(err string) (int, *Response) {
	return 400, &Response{ErrorMessage: err}
}
