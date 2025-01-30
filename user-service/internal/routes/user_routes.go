package routes

import (
	"net/http"

	"user-service/internal/controller"
)

func UserRoutes() {
	http.HandleFunc("/user/register", controller.RegisterHandler)
}
