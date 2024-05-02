package models

type User struct {
	Id       int    `json:"uid,omitempty" gorm:"primary_key"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Admin    bool   `json:"admin,omitempty"`
	Password string `json:"password,omitempty"`
}
