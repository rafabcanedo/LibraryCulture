package router

import (
	"github.com/rafabcanedo/LibraryCulture/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/book", handler.ShowBookHandler)

		v1.POST("/book", handler.CreateBookHanlder)

		v1.PUT("/book", handler.UpdateBookHandler)

		v1.DELETE("/book", handler.DeleteBookHandler)

		v1.GET("/books", handler.ListBooksHandler)

		v1.GET("/user", handler.ShowUserHandler)

		v1.POST("/user", handler.CreateUserHandler)

		v1.PUT("/user", handler.UpdateUserHandler)

		v1.DELETE("/user", handler.DeleteUserHandler)

		v1.GET("/users", handler.ListUsersHandler)

	}
}
