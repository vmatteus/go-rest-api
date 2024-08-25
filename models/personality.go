package models

type Personality struct {
	ID      int    `json:"id"`
	Name    string `json:"name" gorm:"column:nome"`
	History string `json:"history" gorm:"column:historia"`
	Teste   string `json:"teste" gorm:"column:teste"`
}

var Personalities []Personality

func (Personality) TableName() string {
	return "personalidades"
}
