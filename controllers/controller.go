package controllers

import (
	"fmt"
	"net/http"

	"github.com/William-Libero/go-rest-api-com-gin/database"
	"github.com/William-Libero/go-rest-api-com-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func ExibeAluno(c *gin.Context) {
	var aluno models.Aluno
	nome := c.Params.ByName("nome")
	database.DB.First(&aluno, nome)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": fmt.Sprintf("Aluno %s não encontrado", nome),
		})
		return
	}
	c.JSON(200, aluno)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	alunoId := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Find(&aluno, alunoId)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	database.DB.Delete(&aluno)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func AlteraAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")
	database.DB.Where(models.Aluno{CPF: cpf}).First(&aluno)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
