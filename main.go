package main

import (
	"github.com/William-Libero/go-rest-api-com-gin/database"
	"github.com/William-Libero/go-rest-api-com-gin/models"
	"github.com/William-Libero/go-rest-api-com-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "William", RG: "123456", CPF: "123456789"},
		{Nome: "Leo", RG: "654321", CPF: "987654321"},
	}
	routes.HandleRequests()
}
