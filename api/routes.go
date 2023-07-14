package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/domains", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	return router
}
