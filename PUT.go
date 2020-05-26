package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func editUser(c *gin.Context) {
	userid := checkAuth(c)
	var err error
	var good bool
	token := c.Request.Header.Get("Authorization")
	//var data []*User

	if userid == 0 {
		return
	}
	c.Request.ParseForm()
	name := strings.Join(c.Request.PostForm["name"], "")
	surname := strings.Join(c.Request.PostForm["surname"], "")
	pic := strings.Join(c.Request.PostForm["pic"], "")
	birth := strings.Join(c.Request.PostForm["birth"], "")
	mail := strings.Join(c.Request.PostForm["mail"], "")
	password := strings.Join(c.Request.PostForm["password"], "")
	newPassword := strings.Join(c.Request.PostForm["newPassword"], "")

	//common infos change
	if name != "" && surname != "" && birth != "" && mail != "" {
		err, good = editUserData(userid, name, surname, birth, mail, pic)
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
				err = editUserPass(userid, newPassword, token)
			}
		} else {
			c.JSON(403, "Incorrect password")
		}
	}
	if errorReq(c, err) != true {
		if good == false {
			c.JSON(400, "Mail already exists")
		} else {
			c.JSON(200, token)
		}
	} else {
		if good == false {
			c.JSON(400, "Mail already exists")
		} else {
			c.JSON(400, "An error occured")
		}
	}
}
