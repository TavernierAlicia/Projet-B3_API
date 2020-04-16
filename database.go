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

func TestDB() (data []*TestDb) {
	db, dbname := RunDb()
	data = []*TestDb{}
	//// Providers ////
	err := db.Select(&data, testquery)
	if err != nil {
		log.Error("failed to request database ", zap.String("database", err.Error()),
			zap.String("query_name", "testquery"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "testquery"))
	}
	return data
}

func testPost(name string, description string, price float64) (response bool) {
	db, dbname := RunDb()
	//data = []*TestDb{}
	//// Providers ////

	response = false
	db.MustExec(testPostQuery, name, description, price)
	response = true
	//err := db.Exec(name, description, price, testPostQuery)
	if response == false {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "testquery"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "testquery"))
	}
	return response
}

func testUpdate(id int64, name string, description string, price float64) (response bool) {
	db, dbname := RunDb()

	response = false
	db.MustExec(testUpdateQuery, name, description, price, id)
	response = true
	if response == false {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "testquery"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "testquery"))
	}
	return response
}

func testDelete(id int64) (response bool) {
	db, dbname := RunDb()

	response = false
	db.MustExec(testDeleteQuery, id)
	response = true
	if response == false {
		log.Error("failed to request database ", zap.String("database", dbname),
			zap.String("query_name", "testquery"))
	} else {
		log.Info("Request Succeed ", zap.String("database", dbname),
			zap.String("query_name", "testquery"))
	}
	return response
}
