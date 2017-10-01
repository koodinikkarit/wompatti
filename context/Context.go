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

func (c *Context) CreateComputer(
	name string,
) *WompattiModels.Computer {
	computer := &WompattiModels.Computer{
		Name: name,
	}
	c.db.Create(&computer)
	return computer
}

func (c *Context) UpdateComputer(
	computerID uint32,
	name string,
	wolInterfaceID uint32,
	ip string,
	mac string,
) *WompattiModels.Computer {
	var computer WompattiModels.Computer
	c.db.First(&computer, computerID)
	if computer.ID > 0 {
		if name != "" {
			computer.Name = name
		}
		if wolInterfaceID > 0 {
			computer.WolInterfaceID = wolInterfaceID
		}
		if ip != "" {
			computer.IP = ip
		}
		if mac != "" {
			computer.Mac = mac
		}
		c.db.Save(&computer)
	}
	return &computer
}

func (c *Context) RemoveComputer(computerID uint32) bool {
	var computer WompattiModels.Computer
	c.db.First(&computer, computerID)
	if computer.ID > 0 {
		c.db.Delete(&computer)
		return true
	}
	return false
}

func (c *Context) CreateWolInterface(
	ip string,
	port uint32,
) *WompattiModels.WolInterface {
	wolInterface := &WompattiModels.WolInterface{
		IP:   ip,
		Port: port,
	}
	c.db.Create(&wolInterface)
	return wolInterface
}

func (c *Context) UpdateWolInterface(
	ID uint32,
	ip string,
	port uint32,
) *WompattiModels.WolInterface {
	var wolInterface WompattiModels.WolInterface
	c.db.First(&wolInterface, ID)
	if wolInterface.ID > 0 {
		wolInterface.IP = ip
		wolInterface.Port = port
		c.db.Save(&wolInterface)
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
