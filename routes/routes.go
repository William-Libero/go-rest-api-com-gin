package routes

import (
	"github.com/William-Libero/go-rest-api-com-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/alunos/:nome", controllers.ExibeAluno)
	r.Run()
}
