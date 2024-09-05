package handlersCurso

import (
	"net/http"

	"github.com/gabriel-amat/go_api_school/database"
	"github.com/gabriel-amat/go_api_school/models"
	"github.com/gin-gonic/gin"
)

func GetCursosHandler(ctx *gin.Context) {
	var cursos []models.CursoModel

	res := database.DB.Find(&cursos)

	if res.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": res.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, cursos)
}
