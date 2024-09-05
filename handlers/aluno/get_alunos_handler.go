package handlersAluno

import (
	"net/http"

	"github.com/gabriel-amat/go_api_school/database"
	"github.com/gabriel-amat/go_api_school/models"
	"github.com/gin-gonic/gin"
)

func GetAlunosHandler(ctx *gin.Context) {
	var alunos []models.AlunoModel

	res := database.DB.Find(&alunos)

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, alunos)
}

