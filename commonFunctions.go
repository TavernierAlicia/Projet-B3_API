package main

import (
	"github.com/gin-gonic/gin"
)

//no route availlable
func returnError(c *gin.Context) {
	message := "Nous sommes desoles, une erreur s'est produite"
	//display error page
	data := Error{
		Text: message,
	}
	c.JSON(404, data)
}

//just a test
func testGet(c *gin.Context) {
	message := "Working nice"
	data := Responce{
		Text: message,
		Hello: "Hello you <3",
	}
	c.JSON(200, data)
}