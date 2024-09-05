package models

type AlunoModel struct {
	Codigo uint   `json:"codigo" gorm:"primaryKey"`
	Name   string `json:"name"`
}
