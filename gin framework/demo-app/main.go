package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gobook.io/m/Config"
	"gobook.io/m/Models"
	"gobook.io/m/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
