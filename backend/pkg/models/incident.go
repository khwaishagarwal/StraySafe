package models

type Incident struct {
	Id          int     `json:"id,omitempty" gorm:"primary_key"`
	Uid         int     `json:"uid,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	Image       string  `json:"image,omitempty"`
	Resolved    bool    `json:"resolved,omitempty"`
	ResolverId  int     `json:"resolverId,omitempty" gorm:"column:resolverid"`
}
