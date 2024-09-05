package main

import (
	"github.com/gabriel-amat/go_api_school/database"
	handlersAluno "github.com/gabriel-amat/go_api_school/handlers/aluno"
	handlersCurso "github.com/gabriel-amat/go_api_school/handlers/curso"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//InitDataBase and run migrations with GORM
	database.InitDataBase()
	// Inicializa o router Gin
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.Default())

	//Routes
	//Aluno
	router.GET("/alunos", handlersAluno.GetAlunosHandler)
	router.POST("/alunos", handlersAluno.CreateAlunoHandler)
	router.PATCH("/aluno", handlersAluno.UpdateAlunoHandler)
	//Curso
	router.GET("/cursos", handlersCurso.GetCursosHandler)
	router.GET("/curso/matriculas", handlersCurso.GetCursoMatriculasHandler)
	router.POST("/cursos", handlersCurso.CreateCursoHandler)
	router.PATCH("/curso", handlersCurso.UpdateCursoHandler)
	//Matriculas
	router.GET("/cursos/aluno", handlersCurso.GetCursoAlunosHandler)
	router.POST("/cursos/aluno", handlersCurso.AddAlunoCursoHandler)
	//Start
	router.Run(":8000")
}
