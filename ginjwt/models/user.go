package models

type User struct {
	ID uint64 `json:"id"`
	Username string `json:"username" gorm:"Unique"`
	Password string `json:"password"`
}
