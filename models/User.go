package models

type User struct {
	Id           string `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"not null;unique" json:"username"`
	Email        string `json:"email"`
	Password     string `gorm:"not null" json:"password"`
	RefreshToken string `json:"refresh_token"`
}