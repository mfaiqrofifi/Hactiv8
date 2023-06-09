package router

import (
	"Task_Microservice/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/createBook", controller.CreateBook)
	router.GET("/getBookbyID/:idBook", controller.GetBook)
	router.GET("/getAllBook", controller.GetAllBook)
	router.PUT("/updateBook/:bookID", controller.UpdateBook)
	router.DELETE("/deleteBook/:bookId", controller.DeleteBook)
	return router
}
