package core

import (
	"database/sql"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
)

type MySQLDAO struct {
	db *sql.DB
}

func NewMySQLDAO(database, user, password string) *MySQLDAO {
	dao := new(MySQLDAO)
	var err error

	dao.db, err = sql.Open("mymysql", database+"/"+user+"/"+password)
	if err != nil {
		log.Fatal("Opening myql db failed: ", err)
	}
	return dao
}

func Create(value interface{}) string {
	return ""
}

func Read(id string) interface{} {
	return nil
}

func Update(id string, value interface{}) {

}

func Delete(id string) {

}
