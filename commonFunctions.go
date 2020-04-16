package main

import (
	"fmt"
	"strconv"

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
		Text:  message,
		Hello: "Hello you <3",
	}
	c.JSON(200, data)
}

func testRequest(c *gin.Context) {
	data := TestDB()
	c.JSON(200, data)
}

func TestPost(c *gin.Context) {
	response := testPost("Test Tapas", "Tapas de test", 8.5)
	c.JSON(200, response)
}

func TestPostOptions(c *gin.Context) {
	name := c.Param("name")
	description := c.Param("description")
	price, err := strconv.ParseFloat(c.Param("price"), 64)
	if err != nil {
		fmt.Println("KO")
	} else {
		fmt.Println("OK")
	}
	response := testPost(name, description, price)
	c.JSON(200, response)
}

func TestUpdateOptions(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	name := c.Param("name")
	description := c.Param("description")
	price, err := strconv.ParseFloat(c.Param("price"), 64)
	if err != nil {
		fmt.Println("KO")
	} else {
		fmt.Println("OK")
	}
	response := testUpdate(id, name, description, price)
	c.JSON(200, response)
}

func TestDeleteOptions(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("KO")
	} else {
		fmt.Println("OK")
	}
	response := testDelete(id)
	c.JSON(200, response)
}
