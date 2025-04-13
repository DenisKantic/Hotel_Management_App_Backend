package models

type UserEmployee struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Role     string `gorm:"default:'receptionist'"`
}
