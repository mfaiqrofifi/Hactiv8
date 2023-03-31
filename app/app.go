package app

import (
	"DDD/config"
	"DDD/repository"
	"DDD/route"
	"DDD/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)
	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
