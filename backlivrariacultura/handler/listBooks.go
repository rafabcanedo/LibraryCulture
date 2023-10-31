package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBooksHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Books",
	})
}
