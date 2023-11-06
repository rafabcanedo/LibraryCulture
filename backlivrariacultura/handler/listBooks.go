package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/LibraryCulture/schemas"
)

func ListBooksHandler(ctx *gin.Context) {
	books := []schemas.Book{}

	if err := db.Find(&books).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listining books")
		return
	}

	sendSuccess(ctx, "list-books", books)
}
