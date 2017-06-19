package wompatti

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func CreateDb(
	dbUser string,
	dbPass string,
	dbIp string,
	dbPort string,
	dbName string,
) *gorm.DB {
	db, err := gorm.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbIp+":"+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println("Creating database connection failed", err)
	}

	db.AutoMigrate(&Arttu{})
	db.AutoMigrate(&Computer{})
	db.AutoMigrate(&WolDevice{})
	db.AutoMigrate(&Interface{})
	db.AutoMigrate(&DeviceInfo{})
	db.AutoMigrate(&KeyValue{})
	db.AutoMigrate(&Device{})

	return db
}
