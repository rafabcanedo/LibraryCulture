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
	}
}
