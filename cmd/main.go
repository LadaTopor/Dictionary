package main

import (
	"dictionary/internal/service"
	"dictionary/pkg/logs"

	"github.com/labstack/echo/v4"
)

func main() {
	// создаем логгер
	logger := logs.NewLogger(false)

	// подключаемся к базе
	db, err := PostgresConnection()
	if err != nil {
		logger.Fatal(err)
	}

	svc := service.NewService(db, logger)

	svc2 := service.NewService2(db, logger)

	router := echo.New()
	// создаем группу api
	api := router.Group("api")

	// прописываем пути
	api.GET("/word/:id", svc.GetWordById)
	api.POST("/words", svc.CreateWords)
	api.PUT("/word/:id", svc.UpdateWord)
	api.DELETE("/delete/:id", svc.DeleteWord)

	// пути для репортов
	api.GET("/report/:id", svc2.GetReportById)
	api.POST("/reports", svc2.CreateReports)
	api.PUT("/report/:id", svc2.UpdateReportById)
	api.DELETE("/delete/:id", svc2.DeleteReportById)

	// запускаем сервер, чтобы слушал 8000 порт
	router.Logger.Fatal(router.Start(":8000"))
}
