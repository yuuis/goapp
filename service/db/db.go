package db

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "../common"
  "../../config"
)

var db *sql.DB

func DbInit() (*sql.DB, error){
  var err error;
  db, err = sql.Open(config.DBTYPE,""+config.USER+":"+config.PASSWORD+"@tcp("+config.HOSTNAME+":"+config.HOSTPORT+")/"+config.DBNAME+"?parseTime=true")
  db.SetMaxIdleConns(config.DBMAXIDOLECONNECTION)
  db.SetMaxOpenConns(config.DBMAXCONNECTION)
  return db, err
}

func DbClose() {
  if db != nil {
    db.Close()
  }
}

func DbConn() *sql.DB {
  if db == nil {
    _, err := DbInit()
    if err != nil {
      common.WriteErrorLog(config.DEBUG, err, nil)
      panic(err)
    }
  }
  return db
}
