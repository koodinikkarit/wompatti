package wompatti

import (
	"fmt"

	"google.golang.org/grpc"

	_ "github.com/go-sql-driver/mysql"
	gorm "github.com/jinzhu/gorm"
)

// Context struct for context
type Context struct {
	db     *gorm.DB
	server *grpc.Server
}

// GetDb gets db created for context
func (c *Context) GetDb() *gorm.DB {
	return c.db
}

// Starts server process
func (c *Context) Start() {

}

// Stop server
func (c *Context) Stop() {

}

// CreateContext creates new context
func CreateContext(
	user string,
	passwd string,
	ip string,
	port string,
	dbname string,
	serverPort string) *Context {

	db, err := gorm.Open("mysql", user+":"+passwd+"@tcp("+ip+":"+port+")/"+dbname)

	if err != nil {
		fmt.Println("Creating connection failed", user+":"+passwd+"@tcp("+ip+":"+port+")/"+dbname)
	}

	db.AutoMigrate(&Computer{})

	server := CreateService(db, serverPort)

	return &Context{db: db, server: server}
}
