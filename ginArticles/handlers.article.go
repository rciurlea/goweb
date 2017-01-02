package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}
