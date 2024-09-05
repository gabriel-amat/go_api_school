package handlersCurso

import (
	"net/http"

	"github.com/gabriel-amat/go_api_school/database"
	"github.com/gabriel-amat/go_api_school/models"
	"github.com/gin-gonic/gin"
)

func GetCursoAlunosHandler(ctx *gin.Context) {
	var matriculas []models.CursoAluno
	var cursos []models.CursoModel
	alunoId := ctx.Query("userId")

	if alunoId != "" {
		//Matrículas do aluno
		if err := database.DB.Where("codigo_aluno = ?", alunoId).Find(&matriculas).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		//IDs dos cursos das matrículas
		var cursoIDs []int
		for _, matriculas := range matriculas {
			cursoIDs = append(cursoIDs, matriculas.CodigoCurso)
		}
		//Cursos com esses ID`s
		if err := database.DB.Where("codigo IN ?", cursoIDs).Find(&cursos).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		// Retornar os cursos
		ctx.JSON(http.StatusOK, cursos)

	} else {
		res := database.DB.Find(&matriculas)

		if res.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
			return
		}

		ctx.JSON(http.StatusOK, matriculas)
	}

}
