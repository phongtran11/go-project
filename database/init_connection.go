package database

import (
	"fmt"

	"github.com/phongtran11/go-project/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TDatabaseStore struct {
	DB *gorm.DB
}

var databaseStore TDatabaseStore

func ConnectDB() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(localhost:3306)/go_project?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                              // default size for string fields
		DisableDatetimePrecision:  true,                                                                             // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                             // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                             // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                            // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	databaseStore.DB = db

	fmt.Print("Connect to database successfully\n")

	// Enable automatic migrations
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return databaseStore.DB
}
