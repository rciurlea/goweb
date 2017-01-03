package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		render(c, gin.H{"title": "Successful registration & Login"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error(),
		})
	}
}

func showLoginPage(c *gin.Context) {

}

func performLogin(c *gin.Context) {

}

func logout(c *gin.Context) {

}
