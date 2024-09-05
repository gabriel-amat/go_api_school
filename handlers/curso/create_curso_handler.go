package handlersCurso

import (
	"net/http"

	"github.com/gabriel-amat/go_api_school/database"
	"github.com/gabriel-amat/go_api_school/models"
	"github.com/gin-gonic/gin"
)

func CreateCursoHandler(c *gin.Context) {
	var curso models.CursoModel
	
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	database.DB.Create(&curso)
}
