package routes

import (
	"github.com/betocalestini/gin-go-api/controllers"
	"github.com/betocalestini/gin-go-api/middlewares"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.Use(middlewares.ContentTypeMiddleware())
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.SelecionaAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCpf)
	r.Run()
}
