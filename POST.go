package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

//create user
func createUser(c *gin.Context) {
	token := createUserToken()
	c.Request.ParseForm()

	name := strings.Join(c.Request.PostForm["name"], " ")
	surname := strings.Join(c.Request.PostForm["surname"], " ")
	mail := strings.Join(c.Request.PostForm["mail"], " ")
	password := encodePw(strings.Join(c.Request.PostForm["password"], " "))
	confirmPassword := encodePw(strings.Join(c.Request.PostForm["confirmPassword"], " "))
	birth := strings.Join(c.Request.PostForm["birth"], " ")
	phone := strings.Join(c.Request.PostForm["phone"], " ")

	if password != confirmPassword {
		c.JSON(401, "Mismatch passwords")
		return
	}
	userCreate(name, surname, mail, password, birth, phone, token)
	c.JSON(200, "Account created")
}

//authentification
func auth(c *gin.Context) {
	c.Request.ParseForm()
	mail := strings.Join(c.Request.PostForm["mail"], " ")
	password := encodePw(strings.Join(c.Request.PostForm["password"], " "))
	token := authentification(mail, password)
	if token != "" {
		c.JSON(200, token)
	} else {
		c.JSON(401, "Login or password wrong")
	}
}
