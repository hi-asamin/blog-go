package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/marc-town/blog-go/web-api/app/interfaces/database"
)

type SqlHandler struct {
	Conn *gorm.DB
}

const (
	Dialect    = "mysql"
	DBUser     = "dev"
	DBPass     = "Password_01"
	DBProtocol = "tcp(127.0.0.1:3306)"
	DBName     = "blog_app_db"
)

func NewSqlHandler() database.SqlHandler {
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	conn, err := gorm.Open(Dialect, connect)
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
