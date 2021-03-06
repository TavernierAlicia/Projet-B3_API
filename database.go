package main

import (
	"fmt"
	"time"

	//os/exec
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var log *zap.Logger

func RunDb() (*sqlx.DB, string) {
	log, _ = zap.NewProduction()

	defer log.Sync()
	//// IMPORT CONFIG ////
	viper.SetConfigName("conf")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Unable to load config file", zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
	}

	//// DEFINE CONFIG VARIABLES FROM JSON FILE ////
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")
	dbname := viper.GetString("database.dbname")

	//// DB CONNECTION ////
	pathSQL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	db, err := sqlx.Connect("mysql", pathSQL)

	request := "connect database"
	printErr(request, err)

	return db, dbname
}

func userCreate(name string, surname string, mail string, password string, birth string, phone string, token string) (err error, good bool) {
	db, _ := RunDb()

	verif := "try"

	err = db.Get(&verif, verifyDouble, mail)

	printErr(verifyDouble, err)

	if verif != mail {
		_, err = db.Exec(createAccount, name, surname, mail, password, birth, phone, token, mail)
		printErr(createAccount, err)
		good = true
	} else {
		good = false
	}
	return err, good
}

func authentification(mail string, password string) string {
	db, _ := RunDb()

	var token string
	err := db.Get(&token, authReq, mail, password)
	printErr(authReq, err)

	return token
}

func getUserid(token string) int64 {
	db, _ := RunDb()

	var userid int64
	err := db.Get(&userid, getUID, token)
	printErr(getUID, err)

	return userid
}

func getEtabs() (data []*Bars, err error) {
	db, _ := RunDb()

	data = []*Bars{}
	err = db.Select(&data, getAllEtabs)
	printErr(getAllEtabs, err)

	return data, err
}

func getEtabsParams(barType string, barPop string, barDist int64, lat float64, long float64) (data []*Bars, err error) {
	db, _ := RunDb()
	dist := 0
	db2, err := db.Beginx()
	printErr("db begin", err)
	request := "SELECT e.id, e.name, e.description, e.type, e.latitude, e.longitude, e.main_pic, e.date, e.subtype, e.street_num, e.street_name, e.city, e.zip, e.happy, e.happy_end "
	if barDist == 0 && barPop == "all" {
		request += " FROM etabs AS e"
	} else {
		dist = 10
		db2.Query("SET @userLat = ?;", lat)
		db2.Query("SET @userLong = ?;", long)
		request += " FROM (SELECT id, name, description, type, latitude, longitude, main_pic, subtype, date, street_num, street_name, city, zip, happy, happy_end, ACOS(COS(RADIANS(latitude)) * COS(RADIANS(@userLat)) * COS(RADIANS(@userLong) - RADIANS(longitude)) + SIN(RADIANS(latitude)) * SIN(RADIANS(@userLat)) ) * 6371 AS distance_km FROM etabs HAVING distance_km < ?) AS e"
	}
	if barPop == "fav" {
		request += " LEFT JOIN (SELECT COUNT(user_id) AS favNum, etab_id FROM favoris GROUP BY etab_id) AS tempFav ON tempFav.etab_id = e.id"
	}
	if barDist == 1 {
		dist = 1
	}
	if barType != "all" {
		request += " WHERE e.subtype = ?"
	}
	if barPop == "new" {
		request += " ORDER BY e.date DESC LIMIT 20"
	}
	if barPop == "fav" {
		request += " ORDER BY favNum DESC LIMIT 20"
	}

	data = []*Bars{}

	if barType != "all" {
		if dist != 0 {
			err = db2.Select(&data, request, dist, barType)
		} else {
			err = db2.Select(&data, request, barType)
		}
	} else if dist != 0 {
		err = db2.Select(&data, request, dist)
	} else {
		err = db2.Select(&data, request)
	}

	db2.Commit()

	printErr(request, err)

	return data, err
}

func favEtabs(userid int64) (data []*BarsInFavs, err error) {
	db, _ := RunDb()

	data = []*BarsInFavs{}
	err = db.Select(&data, getFavs, userid)

	printErr(getFavs, err)

	return data, err
}

func search(keyPhrase string) (data []*Bars, err error) {
	db, _ := RunDb()
	keyPhrase1 := "%" + keyPhrase + "%"
	keyPhrase2 := keyPhrase + "%"
	data = []*Bars{}
	err = db.Select(&data, searchResult, keyPhrase2, keyPhrase1, keyPhrase2, keyPhrase1)

	printErr(searchResult, err)

	return data, err
}

func getUserData(userid int64) (data []*User, err error) {
	db, _ := RunDb()

	data = []*User{}
	err = db.Select(&data, getUser, userid)

	printErr(getUser, err)

	return data, err
}

func editUserData(userid int64, name string, surname string, birth string, phone string, pic string) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(editUserCm, name, surname, birth, phone, pic, userid)
	printErr(editUserCm, err)

	return err
}

func editUserPass(userid int64, newPassword string, token string) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(editUserPwd, newPassword, token, userid)
	printErr(editUserPwd, err)
	return err
}

func ShowBarView(userid int64, etabid int64) (data BarView, err error) {
	db, _ := RunDb()
	data = BarView{}

	//BAR INFOS
	err = db.Select(&data.BarDetails, showBarDetails, userid, etabid)

	printErr(showBarDetails, err)

	//BAR PICS
	err = db.Select(&data.Pictures, showBarPictures, etabid)

	printErr(showBarPictures, err)

	//BAR ITEMS
	err = db.Select(&data.Items, showBarItems, etabid)

	printErr(showBarItems, err)

	return data, err

}

func AddToFavs(userid int64, etabid int64) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(addFavs, userid, etabid, userid, etabid)
	printErr(addFavs, err)
	return err
}

func DeleteFromFavs(userid int64, etabid int64) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(deleteFav, userid, etabid)
	printErr(deleteFav, err)
	return err
}

func Order(userid int64, t TakeOrder) (err error, command_id int64) {
	db, _ := RunDb()

	tx, err := db.Begin()
	printErr("db begin", err)

	res, err := tx.Exec(addOrder, userid, t.Etab_id, t.Instructions, t.Waiting_time, t.Payment, t.Tip)
	printErr(addOrder, err)

	command_id, err = res.LastInsertId()
	for _, item := range t.Items {
		_, err = tx.Exec(addOrderItems, command_id, item, item)
		printErr(addOrderItems, err)
	}

	tx.Exec(calcPrice, command_id, command_id)
	printErr(calcPrice, err)

	err = tx.Commit()
	printErr("tx commit", err)
	return err, command_id
}

func GetOrder(cmdId int64) (totalData []*OneCommand, err error) {
	db, _ := RunDb()

	data := []*CommandDetails{}
	subData := []*CommandItems{}
	totalData = []*OneCommand{}

	err = db.Select(&data, showTheOrder, cmdId)
	printErr(showTheOrder, err)
	for _, item := range data {
		err = db.Select(&subData, showOrdersDetails, item.Id)
		printErr(showOrdersDetails, err)

		cmd := &OneCommand{
			*item,
			subData,
		}
		totalData = append(totalData, cmd)
	}

	return totalData, err
}

func GetOrders(userid int64) (totalData []*Commands, err error) {
	db, _ := RunDb()

	data := []*Command{}
	totalData = []*Commands{}
	var subData []*CommandItems

	err = db.Select(&data, showOrders, userid)
	printErr(showOrders, err)

	for _, item := range data {
		subData = []*CommandItems{}
		err = db.Select(&subData, showOrdersDetails, item.Id)
		printErr(showOrdersDetails, err)
		cmd := &Commands{
			*item,
			subData,
		}
		totalData = append(totalData, cmd)
	}
	return totalData, err
}
