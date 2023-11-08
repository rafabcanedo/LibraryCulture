package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/LibraryCulture/schemas"
)

func ShowUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	user := schemas.User{}

	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	sendSuccess(ctx, "show-user", user)
}
