package main

import (
	_ "fmt"
	_ "io"
	_ "net/http"
	_ "os"
	_ "strings"

	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
)

//main function
func main() {

	//logFile, err := os.Create("history.log")
	//if err != nil {
	//  panic(err)
	//}
	//gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	//Define router
	router := gin.Default()
	router.Use(limit.MaxAllowed(20))

	//noroute
	router.NoRoute(returnError)

	//GET

	//show bars
	router.GET("app/show/", showBars)

	//search bar
	router.GET("app/search/", searchName)

	//show favs
	router.GET("app/favs/", showFavs)

	//userProfile
	router.GET("app/profile/", getUserProfile)

	//show bar data
	router.GET("app/show/:id", getEtabContent)

	//re-order
	router.GET("app/getOrder/:commandid", getOrder)

	//show orders
	router.GET("app/showOrders", getOrders)

	//POST

	//createUser
	router.POST("app/createUser/", createUser)

	//auth
	router.POST("app/auth/", auth)

	//add Fav
	router.POST("app/favs/add/:etabid", addingFavs)

	//take order
	router.POST("app/takeOrder", takeOrder)

	//PUT

	//modifyProfile
	router.PUT("app/profile/edit/", editUser)

	//DELETE

	//delete fav
	router.DELETE("app/favs/delete/:etabid", deletingFavs)

	router.Run(":9999")
}
