package core

import (
	"database/sql"
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

func (m MySQLDAO) Create(value interface{}) string {
	return ""
}

func (m MySQLDAO) Read(id string) interface{} {
	return nil
}

func (m MySQLDAO) Update(id string, value interface{}) {

}

func (m MySQLDAO) Delete(id string) {

}
