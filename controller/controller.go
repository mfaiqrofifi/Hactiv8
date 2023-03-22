package controller

import (
	"Task_Microservice/domain"
	"fmt"
	"net/http"

	"Task_Microservice/service"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var dataBook domain.BookDomain
	if err := ctx.ShouldBindJSON(&dataBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	dataBookService := service.CreateBook(dataBook)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": dataBookService,
	})
}
func GetBook(ctx *gin.Context) {
	idBook := ctx.Param("idBook")
	dataBook, err := service.GetDataBook(idBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status_error":  err,
			"message_error": fmt.Sprintf("data %s tidak ditemukan", idBook),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": dataBook,
	})

}

func GetAllBook(ctx *gin.Context) {
	dataBook, err := service.AllBook()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"status_error":  err,
			"message_error": "data kosong",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": dataBook,
	})
}