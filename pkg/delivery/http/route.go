package http

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rifftyo/presence-be/pkg/delivery/http/handler"
	"github.com/rifftyo/presence-be/pkg/delivery/http/middleware"
)

func SetUpRoutes(
    app *fiber.App,
    userHandler *handler.UserHandler,
    absenceHandler *handler.AbsenceHandler,
) {
    api := app.Group("/api")

    userRoutes := api.Group("/users")
    userRoutes.Post("/register", userHandler.Register)
    userRoutes.Post("/login", userHandler.Login)
    userRoutes.Get("/profile", middleware.JWTMiddleware(os.Getenv("JWT_SECRET")), userHandler.GetProfile)

    absenceRoutes := api.Group("/absence")
    absenceRoutes.Post("/check-in", middleware.JWTMiddleware(os.Getenv("JWT_SECRET")), absenceHandler.CheckIn)
    absenceRoutes.Put("/check-out", middleware.JWTMiddleware(os.Getenv("JWT_SECRET")), absenceHandler.CheckOut)
    absenceRoutes.Get("/history", middleware.JWTMiddleware(os.Getenv("JWT_SECRET")), absenceHandler.GetHistory)
    absenceRoutes.Get("/history/:historyId", middleware.JWTMiddleware(os.Getenv("JWT_SECRET")), absenceHandler.GetHistoryById)
}
