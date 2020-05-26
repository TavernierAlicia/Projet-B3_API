package main

import (
	"fmt"

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

	fmt.Println(name)

	fmt.Println(surname)
	fmt.Println(pic)

	fmt.Println(birth)
	fmt.Println(phone)

	fmt.Println(password)
	fmt.Println(newPassword)

	//common infos change
	if name != "" && surname != "" && birth != "" && phone != "" {
		err = editUserData(userid, name, surname, birth, phone, pic)
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
		c.JSON(200, token)
	} else {
		c.JSON(400, "An error occured")
	}
}
