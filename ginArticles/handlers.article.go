package main

import (
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
		}
	}
}
