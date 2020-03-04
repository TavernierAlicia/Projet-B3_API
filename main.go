package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	""
)

//define variables
var (
	log *zap.Logger
)

//main function
func main() {

	//zap stuff
	log, _ = zap.NewProduction()
	defer log.Sync()

	//Define router
	router := gin.Default()

	//common routes
	router.NoRoute(returnError())

	//// ---------------------------------------- PROFESSIONNAL SIDE ---------------------------------------- ////

	//GET
	router.GET("app/pro/index", indexPro)
	router.GET("app/pro/:pro_id/infos", showInfosPro)
	router.GET("app/pro/bar", showMyBar)
	router.GET("app/pro/bar/infos", showMyBarInfos)
	router.GET("app/pro/bar/menu", showMyBarMenu)

	router.GET("app/pro/pending_commands", showPendingCommands)

	router.GET("app/pro/all_commands/:begin/:end", showAllCommands)
	router.GET("app/pro/all_commands/:pro_id/:begin/:end/:status", showCommands)

	router.GET("app/pro/:pro_id/my_todo_commands", showMycommands)
	router.GET("app/pro/:pro_id/my_ended/commands", showMyEndedCommands)
	
	router.GET("app/pro/:pro_id/stats", showStats)
	router.GET("app/pro/:pro_id/campaigns", showCampaigns)


	//POST
	router.POST("app/pro/connect/:usermail/:password", connectPro)
	router.POST("app/pro/:pro_id/bar/menu/add/:name/:picture_path/:description/:price/:sale/:disponibility", addToMenu)
	router.POST("app/pro/:pro_id/campaigns/:name/:description/:begin/:end/:send_date", startCampaign)
	router.POST("app/pro/:pro_id/compaigns/:id/relaunch/:begin/:end/:send_date", relaunchCampaign)


	//PUT
	router.PUT("app/pro/:pro_id/bar/menu/:item_id/modify", updateMenu)
	router.PUT("app/pro/:pro_id/bar/infos/modify", updateBarInfos)
	router.PUT("app/pro/:pro_id/infos/modify", updateInfosPro)
	router.PUT("app/pro/:pro_id/command_id/update/status/:status", updateCommandStatus)
	router.PUT("app/pro/:pro_id/command_id/take/", takeCommand)


	//DELETE
	router.DELETE("app/pro/pro_id/bar/menu/delete/:item_id", deleteMenuItem)



	//// ---------------------------------------- PARTICULAR SIDE ---------------------------------------- ////
	//GET
	router.GET("app/part/index", indexPart)
	router.GET("app/part/bars/all", showAllBars)
	router.GET("app/part/bars/all/position", showAllBarsByPos)
	router.GET("app/part/:client_id/favs", showFavs)
	router.GET("app/part/bars/:bar_id", showThisBar)
	router.GET("app/part/:client_id/infos", showInfosPart)
	router.GET("app/part/:client_id/commands", showAllCommandsPart)
	router.GET("app/part/:client_id/commands/:command_id", showThisCommandPart)
	router.GET("app/part/:client_id/cart", showCart)

	//POST
	router.POST("app/part/connect/:usermail/:password", connectPart)
	router.POST("app/part/search/:option/:searchfield/results", searchBar)
	router.POST("app/part/bars/:bar_id/add/:item_id", addCart)
	router.POST("app/part/:client_id/cart/confirm", confirmCart)

	//PUT
	router.PUT("app/part/:client_id/infos/update", updatePartInfos)
	router.PUT("app/part/:client_id/cart/update", updateCart)

	//DELETE
	router.DELETE("app/part/:client_id/cart/delete/:item_id", removeCartItem)
	router.DELETE("app/part/:client_id/cart/delete/", resetCart)



	//router.Run(":3000")
	router.Run()

}
