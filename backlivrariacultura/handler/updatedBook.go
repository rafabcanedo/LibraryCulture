package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/LibraryCulture/schemas"
)

func UpdateBookHandler(ctx *gin.Context) {
	request := UpdateBookRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	book := schemas.Book{}

	if err := db.First(&book, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "book not found")
		return
	}

	// Update Book
	if request.Title != "" {
		book.Title = request.Title
	}

	if request.Author != "" {
		book.Author = request.Author
	}

	if request.Description != "" {
		book.Description = request.Description
	}

	if request.Reserved != nil {
		book.Reserved = *request.Reserved
	}

	if request.Price > 0 {
		book.Price = request.Price
	}

	// Save Book Updated
	if err := db.Save(&book).Error; err != nil {
		logger.Errorf("error updating book: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating book")
		return
	}

	sendSuccess(ctx, "update-book", book)
}
