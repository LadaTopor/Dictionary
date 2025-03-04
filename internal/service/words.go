package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetWordById ищем слово по id
// localhost:8000/api/word/:id
func (s *Service) GetWordById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	word, err := repo.RGetWordById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: word})
}

// CreateWords добавляем в базу новые слова
// localhost:8000/api/words
func (s *Service) CreateWords(c echo.Context) error {
	var wordSlice []Word
	err := c.Bind(&wordSlice)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.wordsRepo
	for _, word := range wordSlice {
		err = repo.CreateNewWords(word.Title, word.Translation)
	}
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// EditWord изменяем слово в базе по id или добавляем новое если такого id нет
// localhost:8000/api/word/:id
func (s *Service) EditWord(c echo.Context) error {
	var word Word
	err := c.Bind(&word)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.wordsRepo
	err = repo.EditWord(word.Title, word.Translation, id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// DeleteWord удаляем слово из базы по id
// localhost:8000/api/word/:id
func (s *Service) DeleteWord(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
	}
	repo := s.wordsRepo
	err = repo.DeleteWord(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.String(http.StatusOK, "OK")
}

// SearchSimilarWords ищем 100 похожих слов по ключевому слову
//
//localhost:8000/api/search/ru
func (s *Service) SearchSimilarWords(c echo.Context) error {
	keyword := c.QueryParam("title")
	if keyword == "" {
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.wordsRepo
	wordSlice, err := repo.SearchWordsByKeyword(keyword)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: wordSlice})
}
