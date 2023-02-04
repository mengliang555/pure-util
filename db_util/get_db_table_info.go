package dbutil

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type TBInfo struct {
	TableName string
	Column    []*ColumnInfo
}

type ColumnInfo struct {
	Key   string
	Value reflect.Kind
}

var mysqlDB *sql.DB

func Init(dbConnectInfo string) {
	// db, err := sql.Open("mysql", 	)
	db, err := sql.Open("mysql", dbConnectInfo)

	if err != nil {
		log.Printf("Failed to init the db with error:%s", err.Error())
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Printf("Failed to ping the db with error:%s", err.Error())
		panic(err)
	}
	mysqlDB = db
}

func GetTableStruct(tablename string) {
	val, er := mysqlDB.Query(fmt.Sprintf("desc %s", tablename))
	if er != nil {
		panic(er)
	}
	for val.Next() {
		val.Scan()
		vall, er := val.Columns()
		if er != nil {
			panic(er)
		}
		for _, v := range vall {
			log.Printf("val:%v\n", v)
		}
	}
}
