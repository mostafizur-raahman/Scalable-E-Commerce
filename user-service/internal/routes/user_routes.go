package routes

import (
	"net/http"

	"user-service/internal/controller"
)

func UserRoutes(mux *http.ServeMux, userController *controller.UserController) {
	mux.HandleFunc("/user/register", userController.RegisterHandler)
}
