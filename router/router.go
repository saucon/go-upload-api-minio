// This file is generated using ucnbrew tool.
// Check out for more info "https://github.com/saucon/ucnbrew"
package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

// setup gin's router
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"responseCode": "40404", "responseMessage": "Invalid Path"})
	})

	return router
}

func GroupingRouter(rtr *gin.Engine, handler ...func(c *gin.Context)) *gin.Engine {
	if len(handler) == 0 {
		log.Fatal("handler is 0")
	}

	api := rtr.Group("/v1.0/go-upload-api-minio")
	api.POST("/upload", handler[0])

	return rtr
}
