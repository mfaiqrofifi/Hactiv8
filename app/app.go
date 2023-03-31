package app

import (
	"ShowCase/config"
	"ShowCase/repository"
	"ShowCase/route"
	"ShowCase/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)
	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
