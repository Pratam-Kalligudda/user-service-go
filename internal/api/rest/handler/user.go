package handler

import (
	api "github.com/Pratam-Kalligudda/user-service-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/user-service-go/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func SetupUserHandler(rh *api.RestHandler) {
	// repo := repository.NewUserRepository(rh.DB)
	// svc := service.NewUserService(&repo)
	// // _ := UserHandler{
	// // 	svc: svc,
	// // }

}
