package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var log *zap.Logger

func TestDb(attempts int) {
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
	_, err = sqlx.Connect("mysql", pathSQL)
	if err != nil {
		log.Error("failed to connect database", zap.String("database", dbname),
			zap.Int("attempts left", attempts), zap.Duration("backoff", time.Second))

		//only on server && not yet

		for attempts >= 5 {
			//on server uniquely
			exec.Command("/bin/sh", "-c", "sudo service restart mysqld")
			//wait...
			time.Sleep(10 * time.Second)
			//reconnect...
			attempts = attempts - 1
			TestDb(attempts)

		}
		return

	} else {
		fmt.Println("DB Connected")
	}
	return
}

func TestWeb(attempts int) {
	resp, err := http.Get("http://orderndrink.com/")
	if err != nil {
		log.Error("failed to connect website", zap.String("address ", "http://orderndrink.com/"),
			zap.Int("attempts left", attempts), zap.Duration("backoff", time.Second))
		for attempts >= 5 {
			//on server uniquely
			exec.Command("/bin/sh", "-c", "sudo /home/ec2-user/Projects/OrderNDrink/Projet-B3_Website/websocket/.main &")
			//wait...
			time.Sleep(10 * time.Second)
			//reconnect...
			attempts = attempts - 1
			TestDb(attempts)

		}
		return
	} else {
		fmt.Println("Website Connected")
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	return
}

func TestApi(attempts int) {
	resp, err := http.Get("http://api.orderndrink.com/ping")
	if err != nil {
		log.Error("failed to connect API", zap.String("address ", "http://api.orderndrink.com/ping"),
			zap.Int("attempts left", attempts), zap.Duration("backoff", time.Second))
		for attempts >= 5 {
			//on server uniquely
			exec.Command("/bin/sh", "-c", "sudo /home/ec2-user/Projects/OrderNDrink/Projet-B3_API/.main &")
			//wait...
			time.Sleep(10 * time.Second)
			//reconnect...
			attempts = attempts - 1
			TestDb(attempts)
		}
		return
	} else {
		fmt.Println("API Connected")
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	return
}

func main() {
	running := true
	for running == true {
		attempts := 5
		TestDb(attempts)
		TestWeb(attempts)
		TestApi(attempts)
		time.Sleep(60 * time.Second)
	}
}
