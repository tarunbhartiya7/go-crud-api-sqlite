package models

import (
	"time"

	"example.com/events/db"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int       `json:"userId"`
}

func (e Event) Save() Event {
	now := time.Now()
	userId := 1
	query := `INSERT INTO events (name, description, location, dateTime, userId) 
		VALUES (?, ?, ?, ?, ?) 
		RETURNING id, name, description, location, dateTime, userId`
	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, now, userId).
		Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)
	if err != nil {
		panic("Could not save event: " + err.Error())
	}
	return e
}

func GetAllEvents() []Event {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		panic("Could not get events: " + err.Error())
	}
	defer rows.Close()

	events := []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			panic("Could not scan event: " + err.Error())
		}
		events = append(events, event)
	}
	return events
}

func GetEventById(id int) (Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	if err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId); err != nil {
		return Event{}, err
	}
	return event, nil
}

func UpdateEvent(id int, event Event) (Event, error) {
	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?, userId = ? WHERE id = ?`
	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId, id)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func DeleteEvent(id int) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
