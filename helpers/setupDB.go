package helpers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	USER := "root"
	PASS := "secret"
	HOST := "mysql-development"
	PORT := "3306"
	DBNAME := "testapp"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to DB...")
	}
	return db
}
