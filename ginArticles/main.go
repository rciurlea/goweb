package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	router.Run()
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
