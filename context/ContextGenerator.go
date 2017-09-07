package WompattiContext

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/wompatti/models"
)

type ContextGenerator struct {
	dbUser string
	dbPass string
	dbIP   string
	dbPort string
	dbName string
	dev    bool
	db     *gorm.DB
}

func (cg *ContextGenerator) getDb() *gorm.DB {
	if cg.db == nil {
		db, err := gorm.Open(
			"mysql",
			cg.dbUser+":"+cg.dbPass+"@tcp("+cg.dbIP+":"+cg.dbPort+")/"+cg.dbName,
		)
		if err != nil {
			fmt.Println("Creating database connection failed", err)
		}
		WompattiModels.Migrate(db)
		cg.db = db
	}
	if cg.dev == true {
		return cg.db.Debug()
	}
	return cg.db
}

func (cg *ContextGenerator) NewContext() *Context {
	return &Context{
		db: cg.getDb().Begin(),
	}
}
