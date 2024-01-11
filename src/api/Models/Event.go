package Models

import "time"

type Event struct {
	ID           uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title        string    `gorm:"size:255;not null" json:"title"`
	DateTime     time.Time `gorm:"not null" json:"date_time"`
	Location     string    `gorm:"size:255;not null" json:"location"`
	Description  string    `gorm:"size:255;not null" json:"description"`
	Organizer    User      `gorm:"foreignkey:OrganizerId"`
	Participants []User    `gorm:"many2many:event_participants;"`
}
