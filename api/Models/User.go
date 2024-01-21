package Models

type User struct {
	ID        uint     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string   `gorm:"size:255;not null" json:"first_name"`
	LastName  string   `gorm:"size:255;not null" json:"last_name"`
	Email     string   `gorm:"size:100;not null;unique" json:"email"`
	Password  string   `gorm:"size:100;not null;" json:"password"`
	RoleId    uint     `gorm:"not null" json:"role_id"`
	Role      Role     `gorm:"foreignkey:RoleId"`
	Events    []*Event `gorm:"many2many:event_participants;"`
}
