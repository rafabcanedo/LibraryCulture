package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/LibraryCulture/schemas"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Name:  request.Name,
		Email: request.Email,
		Idade: request.Idade,
		Cargo: request.Cargo,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("Error creating: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating user on database")
		return
	}

	sendSuccess(ctx, "create-user", user)
}
