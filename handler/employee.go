package handler

import (
	"ShowCase/helper"
	"ShowCase/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateBook(c *gin.Context) {
	in := model.BookDomain{}
	err := c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	res, err := h.app.CreateBook(in)
	if err != nil {
		helper.InternalServer(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

func (h Handler) GetBook(c *gin.Context) {
	res, err := h.app.GetBook()
	if err != nil {
		helper.InternalServer(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

func (h Handler) GetBookbyID(c *gin.Context) {
	idBook := c.Param("idBook")
	dt, err := strconv.Atoi(idBook)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	res, err := h.app.GetBookbyID(dt)
	if err != nil {
		helper.NotFound(c, err.Error())
		return
	}
	helper.Ok(c, res)
}

func (h Handler) UpdateBook(c *gin.Context) {
	idBook := c.Param("idBook")
	dt, err := strconv.Atoi(idBook)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	in := model.BookDomain{}
	in.ID = dt
	err = c.BindJSON(&in)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	res, err := h.app.UpdateBook(in)
	if err != nil {
		helper.NotFound(c, err.Error())
	}
	helper.Ok(c, res)
}

func (h Handler) DeleteBook(c *gin.Context) {
	idBook := c.Param("idBook")
	dt, err := strconv.Atoi(idBook)
	if err != nil {
		helper.BadRequest(c, err.Error())
		return
	}
	res, err := h.app.DeleteBook(dt)
	if err != nil {
		helper.NotFound(c, err.Error())
	}
	helper.OkWithMessage(c, res)
}
