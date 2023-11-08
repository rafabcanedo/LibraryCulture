package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/LibraryCulture/schemas"
)

func CreateBookHanlder(ctx *gin.Context) {
	request := CreateBookRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	book := schemas.Book{
		Title:       request.Title,
		Author:      request.Author,
		Description: request.Description,
		Reserved:    *request.Reserved,
		Price:       request.Price,
	}

	if err := db.Create(&book).Error; err != nil {
		logger.Errorf("Error creating Book: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating book on database")
		return
	}

	sendSuccess(ctx, "create-book", book)
}
