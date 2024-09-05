package handlersCurso

import (
	"net/http"

	"github.com/gabriel-amat/go_api_school/database"
	"github.com/gabriel-amat/go_api_school/models"

	"github.com/gin-gonic/gin"
)

func UpdateCursoHandler(c *gin.Context) {
	var model models.CursoModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	database.DB.Save(&model)
}
