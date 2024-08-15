package models

// @name User
type User struct {
	UUIDBaseModel
	Username string `json:"username" gorm:"uniqueIndex"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

type UserView struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
