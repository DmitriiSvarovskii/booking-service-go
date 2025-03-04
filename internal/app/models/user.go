package models

type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ChatID    int64  `json:"chat_id"`
}
