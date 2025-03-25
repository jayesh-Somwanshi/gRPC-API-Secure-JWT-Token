package model

type Employee struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}
