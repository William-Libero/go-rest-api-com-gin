package controllers

import (
	"github.com/William-Libero/go-rest-api-com-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func ExibeAluno(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Api diz": "Eaí " + nome + " tudo beleza?",
	})
}
