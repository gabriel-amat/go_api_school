package handlersCurso

import (
	"net/http"

	"github.com/gabriel-amat/go_api_school/database"
	"github.com/gabriel-amat/go_api_school/models"
	"github.com/gin-gonic/gin"
)

func GetCursoMatriculasHandler(ctx *gin.Context) {
	var count int64
	cursoID := ctx.Query("cursoId")

	if cursoID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Codigo do curso é obrigatorio!"})
		return
	}

	// Numero de matrículas neste curso
	if err := database.DB.Model(&models.CursoAluno{}).Where("codigo_curso = ?", cursoID).Count(&count).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"alunos_matriculados": count,
	})
}
