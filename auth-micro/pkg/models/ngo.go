package models

type Ngo struct {
	Id        int     `json:"id,omitempty" gorm:"primary_key"`
	Name      string  `json:"name,omitempty"`
	Email     string  `json:"email,omitempty"`
	About     string  `json:"about,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	Password  string  `json:"password,omitempty"`
}
