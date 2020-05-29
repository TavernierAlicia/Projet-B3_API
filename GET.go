package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func showBars(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	//get params
	lat, err := strconv.ParseFloat(c.DefaultQuery("lat", ""), 64)
	long, err := strconv.ParseFloat(c.DefaultQuery("long", ""), 64)
	barType := c.DefaultQuery("type", "all")
	barPop := c.DefaultQuery("popularity", "all")
	barDist, err := strconv.ParseInt(c.DefaultQuery("distance", ""), 10, 64)

	var data []*Bars

	if barType == "all" && barPop == "all" && barDist == 0 {
		data, err = getEtabs()
	} else {
		data, err = getEtabsParams(barType, barPop, barDist, lat, long)
	}
	if errorReq(c, err) != true {
		c.JSON(200, data)
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}

}

func searchName(c *gin.Context) {
	var data []*Bars
	var err error
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	keyPhrase := c.DefaultQuery("search", "")

	if keyPhrase == "" {
		data, err = getEtabs()
	} else {
		data, err = search(keyPhrase)
	}

	if errorReq(c, err) != true {
		c.JSON(200, data)
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}
}

func showFavs(c *gin.Context) {
	var err error
	userid := checkAuth(c)
	if userid == 0 {
		return
	}
	var data []*BarsInFavs

	data, err = favEtabs(userid)
	if errorReq(c, err) != true {
		c.JSON(200, data)
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}

}

func getUserProfile(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}
	var err error
	var data []*User
	data, err = getUserData(userid)
	if errorReq(c, err) != true {
		c.JSON(200, data)
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}

}

func getEtabContent(c *gin.Context) {
	var data BarView
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	barid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	printErr("strconv id", err)
	data, err = ShowBarView(userid, barid)
	if errorReq(c, err) != true {
		c.JSON(200, data)
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}

}

func getOrder(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	cmdId, err := strconv.ParseInt(c.Param("commandid"), 10, 64)
	printErr("strconv comandid", err)

	data, err := GetOrder(cmdId)
	if errorReq(c, err) != true {
		c.JSON(200, data)
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}

}

func getOrders(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	data, err := GetOrders(userid)
	if errorReq(c, err) != true {
		c.JSON(200, data)
	} else {
		c.JSON(400, gin.H{
			"code":    5,
			"message": string("An error occured")})
	}
}
