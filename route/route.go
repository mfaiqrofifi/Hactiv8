package route

import (
	"DDD/handler"
	"DDD/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewServer(app)
	api := r.Group("/book")
	{
		api.POST("", server.CreateBook)
		api.GET("", server.GetBook)
		api.GET("/:idBook", server.GetBookbyID)
		api.PUT("/:idBook", server.UpdateBook)
		api.DELETE("/:idBook", server.DeleteBook)
	}
}
