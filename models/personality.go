package models

type Personality struct {
	ID      int    `json:"id"`
	Name    string `json:"name" gorm:"column:nome"`
	History string `json:"history" gorm:"column:historia"`
}

func (Personality) TableName() string {
	return "personalidades"
}
