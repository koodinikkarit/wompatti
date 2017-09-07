package WompattiContext

import (
	"github.com/jinzhu/gorm"
)

type WompattiDB struct {
	db *gorm.DB
}

func (wdb *WompattiDB) Find(out interface{}, where ...interface{}) *WompattiDB {
	return &WompattiDB{
		db: wdb.db.Find(out, where),
	}
}

func (wdb *WompattiDB) Where(query interface{}, args ...interface{}) *WompattiDB {
	return &WompattiDB{
		db: wdb.db.Where(query, args),
	}
}
func (wdb *WompattiDB) Table(name string) *WompattiDB {
	return &WompattiDB{
		db: wdb.db.Table(name),
	}
}

func (wdb *WompattiDB) Joins(query string, args ...interface{}) *WompattiDB {
	return &WompattiDB{
		db: wdb.db.Joins(query, args),
	}
}

func (wdb *WompattiDB) Or(query interface{}, args ...interface{}) *WompattiDB {
	return &WompattiDB{
		db: wdb.db.Or(query, args),
	}
}
