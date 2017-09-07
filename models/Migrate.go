package WompattiModels

import "github.com/jinzhu/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Command{})
	db.AutoMigrate(&Device{})
	db.AutoMigrate(&DeviceType{})
	db.AutoMigrate(&Keijo{})
	db.AutoMigrate(&Severi{})
	db.AutoMigrate(&TelnetInterface{})
	db.AutoMigrate(&WolInterface{})
}
