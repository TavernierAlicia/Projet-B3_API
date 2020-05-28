package main

import (
	"github.com/gin-gonic/gin"
)

func editUser(c *gin.Context) {
	userid := checkAuth(c)
	var err error

	token := c.Request.Header.Get("Authorization")
	//var data []*User

	if userid == 0 {
		return
	}

	var t UserEdit

	c.BindJSON(&t)

	name := t.Name
	surname := t.Surname
	pic := t.Pic
	birth := t.Birth
	mail := t.Mail
	phone := t.Phone
	password := t.Pass
	newPassword := t.NewPass

	//common infos change
	if name != "" && surname != "" && birth != "" && phone != "" {
		err = editUserData(userid, name, surname, birth, phone, pic)
	} else {
		c.JSON(400, gin.H{
			"code":    6,
			"message": string("Field(s) missing")})
		return
	}

	//password change
	if newPassword != "" {
		if password == "" {
			c.JSON(403, gin.H{
				"code":    7,
				"message": string("No password")})
		}
		newPassword = newPassword
		password = password
		auth := authentification(mail, password)
		if auth != "" {
			if password == newPassword {
				c.JSON(400, gin.H{
					"code":    8,
					"message": string("Same passwords")})
				return
			} else {
				token = createUserToken()
				err = editUserPass(userid, newPassword, token)
			}
		} else {
			c.JSON(403, gin.H{
				"code":    9,
				"message": string("Incorrect password")})
		}
	}
	if errorReq(c, err) != true {
		c.JSON(200, gin.H{
			"code":    0,
			"message": string(token)})
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}
}
