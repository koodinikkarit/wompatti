package WompattiContext

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/wompatti/models"
)

type Context struct {
	db *gorm.DB
}

func (c *Context) Commit() {
	c.db.Commit()
}

func (c *Context) Rollback() {
	c.db.Rollback()
}

func (c *Context) WompattiDB() *WompattiDB {
	return &WompattiDB{
		db: c.db,
	}
}

func (c *Context) GetDb() *gorm.DB {
	return c.db
}

func (c *Context) CreateWolInterface(
	EthernetInterfaceID uint32,
	Mac string,
) *WompattiModels.WolInterface {
	wolInterface := &WompattiModels.WolInterface{
		EthernetInterfaceID: EthernetInterfaceID,
		Mac:                 Mac,
	}
	c.db.Create(&wolInterface)
	return wolInterface
}

func (c *Context) EditWolInterface(
	ID uint32,
	EthernetInterfaceID uint32,
	Mac string,
) *WompattiModels.WolInterface {
	var wolInterface WompattiModels.WolInterface
	c.db.First(&wolInterface, ID)
	if wolInterface.ID > 0 {
		c.db.Updates(wolInterface)
	}
	return &wolInterface
}

func (c *Context) RemoveWolInterface(
	ID uint32,
) bool {
	var wolInterface WompattiModels.WolInterface
	c.db.First(&wolInterface, ID)
	if wolInterface.ID > 0 {
		c.db.Delete(wolInterface)
		return true
	}
	return false
}
