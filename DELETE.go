package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deletingFavs(c *gin.Context) {
	userid := checkAuth(c)
	if userid == 0 {
		return
	}
	etabid, err := strconv.ParseInt(c.Param("etabid"), 10, 64)
	fmt.Println(err)

	DeleteFromFavs(userid, etabid)
}
