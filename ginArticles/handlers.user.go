package main

import "gopkg.in/gin-gonic/gin.v1"

func showRegistrationPage(c *gin.Context) {
	render(c, nil, "register.html")
}

func register(c *gin.Context) {

}
