package route

import (
	"ShowCase/handler"
	"ShowCase/service"

	_ "ShowCase/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
