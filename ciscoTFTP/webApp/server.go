package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

var dbUser string
var dbPass string
var dbName string
var dbhost string
var dbPort int

// func init() {
// 	flag.StringVar(&dbUser, "usr", "dbadmin", "Mongo db user")
// 	flag.StringVar(&dbPass, "pass", "AllDone07", "Mongo db password")
// 	flag.StringVar(&dbName, "db", "jrsbat", "Name of Mongo DB")
// 	flag.StringVar(&dbhost, "host", "ds061246.mlab.com", "hostname/IP of Mongo db")
// 	flag.IntVar(&dbPort, "port", 61246, "port number to connect to Mongo db")
// }

func main() {
	flag.Parse()
	r := gin.Default()
	addr := ":8080"
	apiHandler := newAPIRes(r)
	var mongoConn MongoConnection
	err := mongoConn.createConnection()
	if err != nil {
		panic(err)
	}
	sess, err := mongoConn.getSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	// mongoDB := sess.DB(mongoConn.dbName)

	us := newUserStorage(sess, mongoConn.dbName)
	usrRes := newUserRes(us)
	usrRes.usrStore.name = usrRes.MyName()
	apiHandler.res[usrRes.MyName()] = usrRes

	apiHandler.setResoreces()
	r.Static("/public", "./public")

	fmt.Printf("running at %s\n", addr)
	// go addStuff(store)
	// usr := User{Name: "WhenAmI"}
	// usr.HexID = bson.NewObjectId()
	// err = mongoDB.C("users").Insert(usr)
	// if err != nil {
	// 	panic(err)
	// }

	r.Run(addr)
}

func checkErr(msg string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%s %v", msg, err))
	}
}
