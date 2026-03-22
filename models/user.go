package models

import "example.com/events/db"

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() (User, error) {
	query := `INSERT INTO users (email, password) VALUES (?, ?) RETURNING id`
	err := db.DB.QueryRow(query, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
