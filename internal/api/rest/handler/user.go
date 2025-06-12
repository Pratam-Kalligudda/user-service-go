package handler

import (
	"net/http"
	"time"

	"github.com/Pratam-Kalligudda/user-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/user-service-go/internal/domain"
	"github.com/Pratam-Kalligudda/user-service-go/internal/dto"
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
	pvtRoutes := pubRoutes.Group("/", rh.Auth.Authorize)
	pvtRoutes.Get("/me", handler.GetProfile)
	pvtRoutes.Put("/update", handler.UpdateProfile)
	pvtRoutes.Get("/refresh", handler.Refresh)
	pvtRoutes.Get("/verification", handler.GetVerificationCode)
	pvtRoutes.Post("/verification", handler.VerifyUser)
	pvtRoutes.Post("become-seller", handler.BecomeSeller)
}

func (h *UserHandler) Login(ctx fiber.Ctx) error {
	var user dto.LoginDTO
	if err := ctx.Bind().Body(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	token, refreshToken, err := h.svc.Login(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	h.setCookie(ctx, refreshToken)

	return ctx.JSON(&fiber.Map{
		"message": "succesfully logged in",
		"token":   token,
	})
}

func (h *UserHandler) setCookie(ctx fiber.Ctx, refresh string) {
	cookie := new(fiber.Cookie)
	cookie.Name = "refresh-token"
	cookie.Value = refresh
	cookie.Expires = time.Now().Add(time.Hour * 24 * 7)
	cookie.Secure = false
	ctx.Cookie(cookie)
}

func (h *UserHandler) Register(ctx fiber.Ctx) error {
	var user dto.SignupDTO

	if err := ctx.Bind().Body(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, refreshToken, err := h.svc.Register(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	h.setCookie(ctx, refreshToken)

	return ctx.JSON(&fiber.Map{
		"message": "succesfully registered",
		"token":   token,
	})
}

func (h *UserHandler) UpdateProfile(ctx fiber.Ctx) error {
	var user domain.User

	if err := ctx.Bind().Body(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := h.svc.UpdateUser(user)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "succesfully updated profile",
	})
}

func (h *UserHandler) Refresh(ctx fiber.Ctx) error {
	user, err := h.svc.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := h.svc.Refresh(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "succesfully refreshed token",
		"token":   token,
	})
}

func (h *UserHandler) GetProfile(ctx fiber.Ctx) error {
	email := ctx.Locals("email").(string)
	user, err := h.svc.Repo.FindUserByEmail(email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(user)
}

func (h *UserHandler) GetVerificationCode(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "succesfully sent verification code",
	})
}

func (h *UserHandler) BecomeSeller(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "succesfully became seller",
	})
}

func (h *UserHandler) VerifyUser(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "succesfully verified user",
	})
}
