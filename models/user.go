package models

import (
	"errors"
	"log"

	"example.com/events/db"
	"example.com/events/utils"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() (User, error) {
	query := `INSERT INTO users (email, password) VALUES (?, ?) RETURNING id`
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return User{}, err
	}
	err = db.DB.QueryRow(query, u.Email, hashedPassword).Scan(&u.ID)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (u User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	var hashedPassword string
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &hashedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}
	if err := utils.VerifyPassword(u.Password, hashedPassword); err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
