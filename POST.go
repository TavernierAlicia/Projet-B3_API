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
	password := t.Pass
	confirmPassword := t.ConfirmPass
	birth := t.Birth
	phone := t.Phone

	if password != confirmPassword {
		c.JSON(401, gin.H{
			"code":    4,
			"message": string("Mismatch passwords")})
		return
	} else {
		err, good = userCreate(name, surname, mail, password, birth, phone, token)
		if errorReq(c, err) == false {
			if good == false {
				c.JSON(400, gin.H{
					"code":    3,
					"message": string("Mail already exists")})
			} else {
				c.JSON(200, gin.H{
					"code":    0,
					"message": string("Account created")})
			}
		} else {
			if good == false {
				c.JSON(400, gin.H{
					"code":    3,
					"message": string("Mail already exists")})
			} else {
				c.JSON(400, gin.H{
					"code":    5,
					"message": string("An error occured")})
			}
		}
	}
}

//authentification
func auth(c *gin.Context) {
	var t Auth

	c.BindJSON(&t)

	mail := t.Mail

	password := t.Pass

	token := authentification(mail, password)

	if token != "" {
		c.JSON(200, gin.H{
			"code":    0,
			"message": string(token)})
	} else {
		c.JSON(401, gin.H{
			"code":    2,
			"message": string("Login or password wrong")})
	}
}

func addingFavs(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}
	etabid, err := strconv.ParseInt(c.Param("etabid"), 10, 64)

	printErr("strconv etabid", err)

	err = AddToFavs(userid, etabid)
	if errorReq(c, err) != true {
		c.JSON(200, gin.H{
			"code":    0,
			"message": string("Fav added")})
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}
}

func takeOrder(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	var t TakeOrder

	c.BindJSON(&t)

	err, command_id := Order(userid, t)
	if errorReq(c, err) != true {
		c.JSON(200, gin.H{
			"code":    0,
			"message": command_id})
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}
}
