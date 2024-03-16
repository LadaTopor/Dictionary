package service

import (
	"net/http"
	"strconv"

	"dictionary/internal/words"
	"github.com/labstack/echo/v4"
)

func (s *Service) GetWordById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	repo := words.NewRepo(s.db)
	word, err := repo.RGetWordById(id)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.JSON(http.StatusOK, Response{Object: word})
}

func (s *Service) CreateWords(c echo.Context) error {
	var wordSlice []Word
	err := c.Bind(&wordSlice)
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InvalidParams)
	}

	repo := words.NewRepo(s.db)
	for _, word := range wordSlice {
		err = repo.CreateNewWord(word.Title, word.Translation)
	}
	if err != nil {
		s.logger.Error(err)
		return s.NewError(InternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
