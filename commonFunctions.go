package main

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
func TestDeleteHeader(c *gin.Context) {
	//id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	id, err := strconv.ParseInt(c.Request.Header.Get("id"), 10, 64)
	if err != nil {
		returnError(c)
	} else {
		response := testDelete(id)
		fmt.Println(response)
		if response != true {
			returnError(c)
		} else {
			c.JSON(200, response)
		}
	}
}
*/

//c.Request.Header.Get("keyName")

func checkAuth(c *gin.Context) int64 {
	auth := c.Request.Header.Get("Authorization")
	var userid int64
	if auth != "" {
		userid = getUserid(auth)
		if userid == 0 {
			c.JSON(401, "Unauthorized")
		}
	} else {
		c.JSON(401, "Unauthorized")
		userid = 0
	}
	return userid
}

func createUserToken() string {
	userToken := uuid.New().String()
	return userToken
}

//if no page
func returnError(c *gin.Context) {
	message := "Nous sommes desoles, une erreur s'est produite"
	//display error page
	data := Error{
		Text: message,
	}
	c.JSON(404, data)
}

func encodePw(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
