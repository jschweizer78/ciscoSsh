package tiedot

import "github.com/HouzuoGuo/tiedot/db"

var myDBDir = "./tmp/MyDatabase"

// (Create if not exist) open a database
//var myDB *db.DB
var myDB, err = db.OpenDB(myDBDir)

func doesColExsit(name string) bool {
	cols := myDB.AllCols()
	exsits := false
	for _, col := range cols {
		if col == "Users" {
			exsits = true
			return exsits
		}
	}
	return exsits
}

func GetDB() *db.DB {
	return myDB
}
