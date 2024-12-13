package models

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
