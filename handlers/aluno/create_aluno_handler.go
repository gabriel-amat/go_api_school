package handlersAluno

import (
	"net/http"

	"github.com/gabriel-amat/go_api_school/database"
	"github.com/gabriel-amat/go_api_school/models"
	"github.com/gin-gonic/gin"
)

func CreateAlunoHandler(c *gin.Context) {
	var aluno models.AlunoModel
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	database.DB.Create(&aluno)
}
