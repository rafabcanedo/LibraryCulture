package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/LibraryCulture/schemas"
)

func UpdateUserHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}

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

	user := schemas.User{}

	if err := db.First(&user, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	// Update User
	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Idade > 0 {
		user.Idade = request.Idade
	}

	if request.Cargo != "" {
		user.Cargo = request.Cargo
	}

	// Save User Updated
	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, "update-user", user)
}
