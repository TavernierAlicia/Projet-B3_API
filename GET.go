package main

import (
	"fmt"
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

	fmt.Println(err)

	if barType == "all" && barPop == "all" && barDist == 0 {
		data = getEtabs()
	} else {
		data = getEtabsParams(barType, barPop, barDist, lat, long)
	}
	c.JSON(200, data)
}

func searchName(c *gin.Context) {
	var data []*Bars
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	keyPhrase := c.DefaultQuery("search", "")

	if keyPhrase == "" {
		data = getEtabs()
	} else {
		data = search(keyPhrase)
	}

	c.JSON(200, data)
}

func showFavs(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}
	var data []*BarsInFavs

	data = favEtabs(userid)
	c.JSON(200, data)
}

func getUserProfile(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}
	var data []*User
	data = getUserData(userid)
	c.JSON(200, data)
}

func getEtabContent(c *gin.Context) {
	var data BarView
	userid := checkAuth(c)
	if userid == 0 {
		return
	}

	barid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Println(err)

	data = ShowBarView(userid, barid)
	c.JSON(200, data)
}
