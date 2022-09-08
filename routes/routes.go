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
	r.GET("/:nome", controllers.ExibeAluno)
	r.Run()
}
