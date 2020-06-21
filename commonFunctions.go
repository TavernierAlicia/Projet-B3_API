package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

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

//if no route
func returnError(c *gin.Context) {
	c.JSON(404, gin.H{
		"code":    1,
		"message": string("This route doesn't exists")})
}

//Now encoding with app
/*
func encodePw(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
*/

func errorReq(c *gin.Context, err error) bool {
	if err != nil {
		return true
	} else {
		return false
	}
}

func printErr(request string, err error) {
	if err != nil {
		if request == "connect database" {
			//display error
			log.Error("Connect database Failed", zap.String("Failed", request),
				zap.Error(err))

			//give error to app
			c.JSON(404, gin.H{
				"code":    10,
				"message": string("Unable to connect database")})

			//reboot mysqld
			exec.Command("/bin/sh", "-c", "sudo service restart mysqld")

		} else {
			log.Error("Request failed", zap.String("Request", request),
				zap.Error(err))
		}

	} else {
		if request == "connect database" {
			log.Info("Connexion etablished ", zap.String("database", "orderndrink"),
				zap.Error(err))
		}
	}
	return
}
