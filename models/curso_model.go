package models

type CursoModel struct {
	Codigo    uint   `json:"codigo" gorm:"primaryKey"`
	Descricao string `json:"descricao"`
	Ementa    string `json:"ementa"`
}
