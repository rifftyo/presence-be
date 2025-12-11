package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rifftyo/presence-be/config"
	"github.com/rifftyo/presence-be/internal/delivery/http"
	"github.com/rifftyo/presence-be/internal/delivery/http/handler"
	"github.com/rifftyo/presence-be/internal/repository"
	"github.com/rifftyo/presence-be/internal/usecase"
)

func main() {
	db, _ := config.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)
	
	absenceRepo := repository.NewAbsenceRepository(db)
	absenceUseCase := usecase.NewAbsenceUseCase(absenceRepo)
	absenceHandler := handler.NewAbsenceHandler(absenceUseCase)

	app := fiber.New()
	app.Static("/uploads", "./uploads")
	
	http.SetUpRoutes(app, userHandler, absenceHandler)

	app.Listen(":3000")
}