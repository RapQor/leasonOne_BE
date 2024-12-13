package main

import (
	"app/internal/handlers"
	"app/internal/repositories"
	"app/internal/routes"
	"app/internal/services"
	"app/pkg/db"
	"app/pkg/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Inisialisasi koneksi database
	db.Init()
	defer db.DB.Close()

	// Inisialisasi service
	userRepository := repositories.Repos(db.DB)
	serviceGlobal := services.ServiceGlobal{
		CreateUser:        userRepository,
		CheckExistingUser: userRepository,
		GetUserByID:       userRepository,
		GetUserByUsername: userRepository,
		GetAllUsers:       userRepository,
	}
	handlers.InitHandler(&serviceGlobal)

	// Inisialisasi router
	mux := http.NewServeMux()
	routes.AuthRoutes(mux)

	// Bungkus router dengan middleware CORS
	handlerWithCORS := middleware.CORS(mux)

	// Menjalankan server
	fmt.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
