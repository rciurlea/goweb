package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/", hello)

	router.Run(":3000")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "muie steaua!\n")
}
