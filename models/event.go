package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int       `json:"userId"`
}

var events = []Event{
	{
		ID:          1,
		Name:        "Event 1",
		Description: "Description 1",
		Location:    "Location 1",
		DateTime:    time.Now(),
		UserId:      1,
	},
	{
		ID:          2,
		Name:        "Event 2",
		Description: "Description 2",
		Location:    "Location 2",
		DateTime:    time.Now(),
		UserId:      2,
	},
}

func (e Event) Save() Event {
	event := Event{
		ID:          len(events) + 1,
		Name:        e.Name,
		Description: e.Description,
		Location:    e.Location,
		DateTime:    time.Now(),
		UserId:      1,
	}
	events = append(events, event)
	return event
}

func GetAllEvents() []Event {
	return events
}

func GetEventById(id int) Event {
	for _, event := range events {
		if event.ID == id {
			return event
		}
	}
	return Event{}
}
