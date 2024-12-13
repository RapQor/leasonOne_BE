package routes

import (
	"app/internal/handlers"
	"net/http"
)

func AuthRoutes(mux *http.ServeMux) {
	// Register routes
	mux.HandleFunc("/api/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/login", handlers.LoginHandler)
	mux.HandleFunc("/api/current-user", handlers.CurrentUserHandler)
	mux.HandleFunc("/api/all-users", handlers.GetAllUsers)
	mux.HandleFunc("/api/logout", handlers.LogoutHandler)
}
