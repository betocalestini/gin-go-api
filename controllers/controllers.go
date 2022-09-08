package controllers

import (
	"github.com/betocalestini/gin-go-api/database"
	"github.com/betocalestini/gin-go-api/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Mensagem": "E ai " + nome + ", tudo beleza?",
	})
}

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(200, aluno)
}

func SelecionaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID != 0 {
		database.DB.Delete(&aluno, id)
		c.JSON(200, gin.H{
			"data": "Aluno deletado com sucesso",
		})
		return
	}
	c.JSON(400, gin.H{
		"Error": "Algo deu errado",
	})
}

func EditaAluno(c *gin.Context) {

	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(200, aluno)
}

func BuscaAlunoPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(200, aluno)

}
