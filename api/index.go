package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/rifftyo/presence-be/config"

	delivery "github.com/rifftyo/presence-be/internal/delivery/http"

	"github.com/rifftyo/presence-be/internal/delivery/http/handler"
	"github.com/rifftyo/presence-be/internal/repository"
	"github.com/rifftyo/presence-be/internal/usecase"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db, _ := config.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	absenceRepo := repository.NewAbsenceRepository(db)
	absenceUseCase := usecase.NewAbsenceUseCase(absenceRepo)
	absenceHandler := handler.NewAbsenceHandler(absenceUseCase)

	app := fiber.New()
	
	delivery.SetUpRoutes(app, userHandler, absenceHandler)

	adaptor.FiberApp(app)(w, r)
}