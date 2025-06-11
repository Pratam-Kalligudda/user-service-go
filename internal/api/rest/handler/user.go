package handler

import (
	"net/http"

	"github.com/Pratam-Kalligudda/user-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/user-service-go/internal/repository"
	"github.com/Pratam-Kalligudda/user-service-go/internal/service"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserHandler(rh *rest.RestHandler) {
	app := rh.App
	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}
	handler := UserHandler{
		svc: svc,
	}

	pubRoutes := app.Group("/user")
	pubRoutes.Post("/login", handler.Login)
	pubRoutes.Post("/register", handler.Register)
}

func (h *UserHandler) Login(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "succesfully logged in",
	})
}

func (h *UserHandler) Register(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "succesfully registered",
	})
}

func (h *UserHandler) UpdateProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "succesfully updated profile",
	})
}

func (h *UserHandler) Refresh(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "succesfully refreshed token",
	})
}
