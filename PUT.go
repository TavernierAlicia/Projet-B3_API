package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func editUser(c *gin.Context) {
	userid := checkAuth(c)
	token := c.Request.Header.Get("Authorization")
	//var data []*User

	if userid == 0 {
		return
	}
	c.Request.ParseForm()
	name := strings.Join(c.Request.PostForm["name"], "")
	surname := strings.Join(c.Request.PostForm["surname"], "")
	birth := strings.Join(c.Request.PostForm["birth"], "")
	mail := strings.Join(c.Request.PostForm["mail"], "")
	password := strings.Join(c.Request.PostForm["password"], "")
	newPassword := strings.Join(c.Request.PostForm["newPassword"], "")

	//common infos change
	if name != "" && surname != "" && birth != "" && mail != "" {
		editUserData(userid, name, surname, birth, mail)
	} else {
		c.JSON(400, "Field(s) missing")
		return
	}

	//password change
	if newPassword != "" {
		if password == "" {
			c.JSON(403, "No password")
		}
		newPassword = encodePw(newPassword)
		password = encodePw(password)
		auth := authentification(mail, password)
		if auth != "" {
			if password == newPassword {
				c.JSON(400, "Same passwords")
				return
			} else {
				token = createUserToken()
				editUserPass(userid, newPassword, token)
			}
		}
	}
	c.JSON(200, token)
}
