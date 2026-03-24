package routes

import "example.com/events/models"

type signupRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret123"`
}

type signupUserData struct {
	ID    int    `json:"id" example:"1"`
	Email string `json:"email" example:"user@example.com"`
}

type signupResponse struct {
	Message string         `json:"message" example:"User created successfully"`
	User    signupUserData `json:"user"`
}

type loginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret123"`
}

type loginResponse struct {
	Message string `json:"message" example:"Login successful"`
	Token   string `json:"token" example:"eyJhbGciOiJI..."`
}

type createEventRequest struct {
	Name        string `json:"name" example:"Concert"`
	Description string `json:"description" example:"Live music performance"`
	Location    string `json:"location" example:"Central Park"`
}

type updateEventRequest struct {
	Name        string `json:"name" example:"Updated Concert"`
	Description string `json:"description" example:"Updated details"`
	Location    string `json:"location" example:"New Venue"`
}

type eventsListResponse []models.Event

type messageResponse struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Error string `json:"error"`
}
