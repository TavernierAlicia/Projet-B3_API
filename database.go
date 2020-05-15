package main

import (
	"fmt"
	"time"

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
	if err != nil {
		log.Error("failed to connect database", zap.String("database", dbname),
			zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
		return db, dbname

	} else {
		log.Info("Connexion etablished ", zap.String("database", dbname),
			zap.Int("attempt", 3), zap.Duration("backoff", time.Second))
	}
	return db, dbname
}

func userCreate(name string, surname string, mail string, password string, birth string, phone string, token string) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(createAccount, name, surname, mail, password, birth, phone, token)
	return err
}

func authentification(mail string, password string) string {
	db, _ := RunDb()

	var token string
	err := db.Get(&token, authReq, mail, password)
	printErr(err)

	return token
}

func getUserid(token string) int64 {
	db, _ := RunDb()

	var userid int64
	err := db.Get(&userid, getUID, token)
	printErr(err)

	return userid
}

func getEtabs() (data []*Bars, err error) {
	db, _ := RunDb()

	data = []*Bars{}
	err = db.Select(&data, getAllEtabs)
	printErr(err)

	return data, err
}

func getEtabsParams(barType string, barPop string, barDist int64, lat float64, long float64) (data []*Bars, err error) {
	db, _ := RunDb()
	dist := 0
	db2, err := db.Beginx()
	printErr(err)
	request := "SELECT e.id, e.name, e.description, e.type, e.latitude, e.longitude, e.main_pic, e.date, e.subtype"
	if barDist == 0 && barPop == "all" {
		request += " FROM etabs AS e"
	} else {
		dist = 10
		db2.Query("SET @userLat = ?;", lat)
		db2.Query("SET @userLong = ?;", long)
		request += " FROM (SELECT id, name, description, type, latitude, longitude, main_pic, subtype, date, ACOS(COS(RADIANS(latitude)) * COS(RADIANS(@userLat)) * COS(RADIANS(@userLong) - RADIANS(longitude)) + SIN(RADIANS(latitude)) * SIN(RADIANS(@userLat)) ) * 6371 AS distance_km FROM etabs HAVING distance_km < ?) AS e"
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

	printErr(err)

	return data, err
}

func favEtabs(userid int64) (data []*BarsInFavs, err error) {
	db, _ := RunDb()

	data = []*BarsInFavs{}
	err = db.Select(&data, getFavs, userid)

	printErr(err)

	return data, err
}

func search(keyPhrase string) (data []*Bars, err error) {
	db, _ := RunDb()
	keyPhrase1 := "%" + keyPhrase + "%"
	keyPhrase2 := keyPhrase + "%"
	data = []*Bars{}
	err = db.Select(&data, searchResult, keyPhrase2, keyPhrase1, keyPhrase2, keyPhrase1)

	printErr(err)

	return data, err
}

func getUserData(userid int64) (data []*User, err error) {
	db, _ := RunDb()

	data = []*User{}
	err = db.Select(&data, getUser, userid)

	printErr(err)

	return data, err
}

func editUserData(userid int64, name string, surname string, birth string, mail string, pic string) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(editUserCm, name, surname, birth, mail, pic, userid)
	return err
}

func editUserPass(userid int64, newPassword string, token string) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(editUserPwd, newPassword, token, userid)
	return err
}

func ShowBarView(userid int64, etabid int64) (data BarView, err error) {
	db, _ := RunDb()
	data = BarView{}

	//BAR INFOS
	err = db.Select(&data.BarDetails, showBarDetails, userid, etabid)

	printErr(err)

	//BAR PICS
	err = db.Select(&data.Pictures, showBarPictures, etabid)

	printErr(err)

	//BAR ITEMS
	err = db.Select(&data.Items, showBarItems, etabid)

	printErr(err)

	return data, err

}

func AddToFavs(userid int64, etabid int64) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(addFavs, userid, etabid)
	return err
}

func DeleteFromFavs(userid int64, etabid int64) (err error) {
	db, _ := RunDb()

	_, err = db.Exec(deleteFav, userid, etabid)
	return err
}

func Order(userid int64, t TakeOrder) (err error) {
	db, _ := RunDb()

	tx, err := db.Begin()
	printErr(err)

	res, err := tx.Exec(addOrder, userid, t.Etab_id, t.Instructions, t.Waiting_time, t.Payment, t.Tip)
	printErr(err)

	command_id, err := res.LastInsertId()
	for _, item := range t.Items {
		_, err = tx.Exec(addOrderItems, command_id, item, item)
		printErr(err)
	}

	tx.Exec(calcPrice, command_id, command_id)

	err = tx.Commit()

	return err
}

func GetOrder(cmdId int64) (data []*OrderItems, err error) {
	db, _ := RunDb()

	data = []*OrderItems{}

	err = db.Select(&data, reOrder, cmdId)

	printErr(err)

	return data, err
}

func GetOrders(userid int64) (totalData []*Commands, err error) {
	db, _ := RunDb()

	data := []*Command{}
	subData := []*CommandItems{}
	totalData = []*Commands{}

	err = db.Select(&data, showOrders, userid)

	for _, item := range data {

		err = db.Select(&subData, showOrdersDetails, item.Id)
		printErr(err)

		cmd := &Commands{
			*item,
			subData,
		}
		totalData = append(totalData, cmd)
	}

	printErr(err)
	return totalData, err
}
