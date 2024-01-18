package Models

import "time"

type UpdateEventInput struct {
	Title        string    `json:"title"`
	DateTime     time.Time `json:"date_time"`
	Location     string    `json:"location"`
	Description  string    `json:"description"`
	OrganizerId  uint      `json:"organizer_id"`
}