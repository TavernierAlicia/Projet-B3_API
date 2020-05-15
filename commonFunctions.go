package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

func errorReq(c *gin.Context, err error) bool {
	if err != nil {
		return true
	} else {
		return false
	}
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
