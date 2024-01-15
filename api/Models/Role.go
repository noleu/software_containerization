package Models

type Role struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null" json:"name"`
}
