package routes

import (
	"github.com/William-Libero/go-rest-api-com-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/alunos/:nome", controllers.ExibeAluno)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:nome", controllers.DeletaAluno)
	r.PUT("/alunos/:id", controllers.AlteraAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaPorCpf)
	r.Run()
}
