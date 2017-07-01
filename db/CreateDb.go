package wompatti

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectionsQuery(db *gorm.DB, after uint32, before uint32, first uint32, last uint32) (*gorm.DB, uint32) {
	var startCursor uint32
	q := db.Debug().Limit(5000).Offset(after)
	if first > 0 {
		q = q.Limit(first)
	}
	if before > 0 {
		q = q.Limit(before - after)
	}
	if first > 0 && before > after {
		q = q.Offset(before - after - last)
		startCursor = before - after - last
	}
	return q, startCursor
}

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
	db.AutoMigrate(&WolInterface{})
	db.AutoMigrate(&DeviceType{})
	db.AutoMigrate(&Command{})
	db.AutoMigrate(&TelnetInterface{})

	return db
}
