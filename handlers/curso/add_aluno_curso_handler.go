package handlersCurso

import (
	"fmt"
	"net/http"

	"github.com/gabriel-amat/go_api_school/database"
	"github.com/gabriel-amat/go_api_school/models"
	"github.com/gin-gonic/gin"
)

func AddAlunoCursoHandler(c *gin.Context) {
	var model models.CursoAluno
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	// Verificar se o aluno já está matriculado nesse curso
	result := database.DB.Where("codigo_aluno = ? AND codigo_curso = ?", model.CodigoAluno, model.CodigoCurso).First(&model)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "O estudante já esta matriculado neste curso."})
		return
	}

	//Verificar se o curso esta lotado
	var alunos int64
	database.DB.Model(&models.CursoAluno{}).Where("codigo_curso = ?", model.CodigoCurso).Count(&alunos)
	if alunos >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "O curso esta cheio"})
		return
	}

	//Verificar quantidade de curso que o aluno esta matriculado
	var cursos int64
	database.DB.Model(&models.CursoAluno{}).Where("codigo_aluno = ?", model.CodigoAluno).Count(&cursos)
	if cursos >= 3 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "O aluno já está matriculado em 3 cursos"})
		return
	}

	//Matricula o aluno no curso
	if err := database.DB.Create(&model).Error; err != nil {
		msg := fmt.Sprintf("Erro ao matricular aluno: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aluno matriculado com sucesso!"})

}
