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

func userCreate(name string, surname string, mail string, password string, birth string, phone string, token string) (response bool) {
	db, dbname := RunDb()

	response = false
	db.MustExec(createAccount, name, surname, mail, password, birth, phone, token)
	response = true
	if response == false {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "createAccount"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "createAccount"))
	}
	return response
}

func authentification(mail string, password string) string {
	db, dbname := RunDb()

	var token string
	err := db.Get(&token, authReq, mail, password)
	if err == nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "authReq"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "authReq"))
	}
	return token
}

func getUserid(token string) int64 {
	db, dbname := RunDb()

	var userid int64
	err := db.Get(&userid, getUID, token)
	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "getUID"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "getUID"))
	}
	return userid
}

func getEtabs() (data []*Bars) {
	db, dbname := RunDb()

	data = []*Bars{}
	err := db.Select(&data, getAllEtabs)

	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "getAllEtabs"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "getAllEtabs"))
	}
	return data
}

func getEtabsParams(barType string, barPop string, barDist int64, lat float64, long float64) (data []*Bars) {
	db, dbname := RunDb()
	dist := 0
	var err error
	db2, err := db.Beginx()
	request := "SELECT e.id, e.name, e.description, e.type, e.latitude, e.longitude, e.main_pic, e.date, e.subtype"
	if barDist == 0 && barPop == "all" {
		request += " FROM etabs AS e"
	} else {
		dist = 10
		db2.Query("SET @userLat = ?;", lat)
		db2.Query("SET @userLong = ?;", long)
		request += " FROM (SELECT id, name, description, type, latitude, longitude, main_pic, subtype, date, ACOS(COS(RADIANS(latitude)) * COS(RADIANS(@userLat)) * COS(RADIANS(@userLong) - RADIANS(longitude)) + SIN(RADIANS(latitude)) * SIN(RADIANS(@userLat)) ) * 6371 AS distance_km FROM etabs HAVING distance_km < ?) AS e "
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
			fmt.Println(barType)
			fmt.Println(dist)
			fmt.Println(request)
		} else {
			err = db2.Select(&data, request, barType)
		}
	} else if dist != 0 {
		err = db2.Select(&data, request, dist)
	} else {
		err = db2.Select(&data, request)
	}

	db2.Commit()

	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "getAllEtabsParams"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "getAllEtabsParams"))
	}

	return data
}

func favEtabs(userid int64) (data []*Bars) {
	db, dbname := RunDb()

	data = []*Bars{}
	err := db.Select(&data, getFavs, userid)

	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "getFavs"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "getFavs"))
	}
	return data
}

func search(keyPhrase string) (data []*Bars) {
	db, dbname := RunDb()
	keyPhrase1 := "%" + keyPhrase + "%"
	keyPhrase2 := keyPhrase + "%"
	data = []*Bars{}
	err := db.Select(&data, searchResult, keyPhrase2, keyPhrase1, keyPhrase2, keyPhrase1)

	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "searchResult"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "searchResult"))
	}
	return data
}

func getUserData(userid int64) (data []*User) {
	db, dbname := RunDb()

	data = []*User{}
	err := db.Select(&data, getUser, userid)

	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "getUser"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "getUser"))
	}
	return data
}

func editUserData(userid int64, name string, surname string, birth string, mail string) {
	db, dbname := RunDb()

	db.Exec(editUserCm, name, surname, birth, mail, userid)
	fmt.Println(dbname)
	fmt.Println(editUserCm + name)
	return
}

func editUserPass(userid int64, newPassword string, token string) {
	db, dbname := RunDb()

	db.Exec(editUserPwd, newPassword, token, userid)
	fmt.Println(dbname)
	return
}

func ShowBarView(userid int64, etabid int64) (data BarView) {
	db, dbname := RunDb()
	data = BarView{}
	var err error

	//BAR INFOS
	err = db.Select(&data.BarDetails, showBarDetails, etabid)

	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "ShowBarDetails"))
		fmt.Println(err)
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "ShowBarDetails"))
	}

	//BAR PICS
	err = db.Select(&data.Pictures, showBarPictures, etabid)

	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "showBarPictures"))
		fmt.Println(err)
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "showBarPictures"))
	}

	//BAR ITEMS
	err = db.Select(&data.Items, showBarItems, etabid, userid, etabid)

	if err != nil {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "showBarItems"))
		fmt.Println(err)
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "showBarItems"))
	}

	return data

}
