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
			"error_status":  "Data not Found",
			"error_message": fmt.Sprintf("data with id %s not found", idBook),
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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Empthy list",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": dataBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	var dataBook domain.BookDomain
	if errs := ctx.ShouldBindJSON(&dataBook); errs != nil {
		ctx.AbortWithError(http.StatusBadRequest, errs)
		return
	}
	status, err := service.UpdateBook(bookID, dataBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Data with id %s not found", bookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": status,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookId")
	stat, err := service.DeleteBook(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not Found",
			"error_message": fmt.Sprintf("Data with id %s not found", bookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": stat,
	})
}
