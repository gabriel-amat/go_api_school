package models

type CursoAluno struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	Codigo      int        `json:"codigo"`
	CodigoAluno int        `json:"codigoAluno"`
	CodigoCurso int        `json:"codigoCurso"`
}
