package repositories

import (
	"app/internal/models"
	"errors"
)

type UserRepo interface {
	CreateUser(user *models.User) error
	CheckExistingUser(username string) error
	GetUserByID(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
}

func (r *repos) CreateUser(user *models.User) error {
	result, err := r.db.Exec("INSERT INTO users (name, age, username, password) VALUES (?, ?, ?, ?)", user.Name, user.Age, user.Username, user.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = int(id)
	return nil
}

func (r *repos) CheckExistingUser(username string) error {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("pengguna sudah ada")
	}
	return nil
}

func (r *repos) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Age, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repos) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.Id, &user.Name, &user.Age, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repos) GetAllUsers() ([]*models.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Username, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
