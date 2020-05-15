package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func deletingFavs(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}
	etabid, err := strconv.ParseInt(c.Param("etabid"), 10, 64)

	err = DeleteFromFavs(userid, etabid)
	if errorReq(c, err) != true {
		c.JSON(200, "deleted")
	} else {
		c.JSON(400, "An error occured")
	}

}
