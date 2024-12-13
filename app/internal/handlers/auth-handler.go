package handlers

import (
	"app/internal/models"
	"app/internal/services"
	"encoding/json"
	"net/http"
	"strings"
)

type Response struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Username string `json:"username"`
	Password string `json:"password"`
	Message  string `json:"message"`
}

var service *services.ServiceGlobal

// Initialize service in an actual application setup
func InitHandler(s *services.ServiceGlobal) {
	service = s
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode request tidak valid", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	// Call RegisterUser on ServiceGlobal
	if err := service.RegisterUser(&user); err != nil {
		http.Error(w, "Gagal mendaftarkan pengguna", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Response{
		Id:       user.Id,
		Name:     user.Name,
		Age:      user.Age,
		Username: user.Username,
		Password: user.Password,
		Message:  "Berhasil mendaftar!",
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode request tidak valid", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Data tidak valid", http.StatusBadRequest)
		return
	}

	if user.Username == "" {
		http.Error(w, "Username tidak boleh kosong", http.StatusBadRequest)
		return
	}

	// Authenticate the user using ServiceGlobal
	existingUser, err := service.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate the token
	token, err := services.GenerateAuthToken(existingUser)
	if err != nil {
		http.Error(w, "Gagal membuat token", http.StatusInternalServerError)
		return
	}

	// Set token to cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,                    // Prevent access to the cookie via JavaScript
		Secure:   true,                    // Use cookies over HTTPS only (recommend for production)
		SameSite: http.SameSiteStrictMode, // Prevent CSRF
		Path:     "/",                     // Path where the cookie is available
	})

	// Return token and success message in response body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   token, // Include token in the response
	})
}

func CurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode request tidak valid", http.StatusMethodNotAllowed)
		return
	}

	// Get token from cookies
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is required", http.StatusUnauthorized)
		return
	}

	// The token is expected to be in the format "Bearer <token>"
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader { // If the prefix was not found
		http.Error(w, "Token tidak valid", http.StatusUnauthorized)
		return
	}

	// Get the user from the token
	user, err := service.GetUserFromToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Send user data back in response
	json.NewEncoder(w).Encode(Response{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Age:      user.Age,
		Password: user.Password,
		Message:  "Berhasil mengambil data pengguna",
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode request tidak valid", http.StatusMethodNotAllowed)
		return
	}

	users, err := service.FindAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode request tidak valid", http.StatusMethodNotAllowed)
		return
	}

	// Clear the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		HttpOnly: true,                    // Prevent access to the cookie via JavaScript
		Secure:   true,                    // Use cookies over HTTPS only (recommend for production)
		SameSite: http.SameSiteStrictMode, // Prevent CSRF
		Path:     "/",                     // Path where the cookie is available
	})

	// Return a success message
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
