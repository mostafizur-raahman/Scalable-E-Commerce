package routes

import (
	"net/http"

	"user-service/internal/controller"
)

func UserRoutes(mux *http.ServeMux, userController *controller.UserController) {
	mux.HandleFunc("/user/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		userController.RegisterHandler(w, r)
	})

	mux.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		userController.LoginHandler(w, r)
	})
}
