package models

import (
	"fmt"
	"go-company/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"os"
)

//var DB *gorm.DB
type dbHandle struct{
	 *gorm.DB
}
var db *dbHandle


//SetupDatabase  - setup database
func SetupDatabase() {
	
	pqlConn :=  fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", configs.DatabaseHost,configs.DatabaseUser,
		configs.DatabasePassword, configs.DatabaseName, configs.DatabasePort)
	dbHandle1, err := gorm.Open(postgres.Open(pqlConn), &gorm.Config{})
	
	fmt.Println(dbHandle1)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database:  %s",pqlConn))
	}	
	
	db = &dbHandle{ dbHandle1}
	db.AutoMigrate(&Company{})
}

func  GetDbHandle()  *dbHandle{
	return db
}