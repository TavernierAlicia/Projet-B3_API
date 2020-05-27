package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

//create user
func createUser(c *gin.Context) {
	var err error
	var good bool
	token := createUserToken()

	var t UserCreate

	c.BindJSON(&t)

	name := t.Name
	surname := t.Surname
	mail := t.Mail
	password := encodePw(t.Pass)
	confirmPassword := encodePw(t.ConfirmPass)
	birth := t.Birth
	phone := t.Phone

	if password != confirmPassword {
		c.JSON(401, "Mismatch passwords")
		return
	} else {
		err, good = userCreate(name, surname, mail, password, birth, phone, token)
		if errorReq(c, err) == false {
			if good == false {
				c.JSON(400, "Mail already exists")
			} else {
				c.JSON(200, "Account created")
			}
		} else {
			if good == false {
				c.JSON(400, "Mail already exists")
			} else {
				c.JSON(400, "An error occured")
			}
		}
	}
}

//authentification
func auth(c *gin.Context) {
	var t Auth

	c.BindJSON(&t)

	mail := t.Mail

	password := encodePw(t.Pass)

	token := authentification(mail, password)

	if token != "" {
		c.JSON(200, token)
	} else {
		c.JSON(401, "Login or password wrong")
	}
}

func addingFavs(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}
	etabid, err := strconv.ParseInt(c.Param("etabid"), 10, 64)

	printErr(err)

	err = AddToFavs(userid, etabid)
	if errorReq(c, err) != true {
		c.JSON(200, "Added!")
	} else {
		c.JSON(400, "An error occured")
	}
}

func takeOrder(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	var t TakeOrder

	c.BindJSON(&t)

	err := Order(userid, t)
	if errorReq(c, err) != true {
		c.JSON(200, "Ordered!")
	} else {
		c.JSON(400, "An error occured")
	}
}
