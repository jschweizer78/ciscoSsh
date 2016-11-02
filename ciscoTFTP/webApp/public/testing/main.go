package main

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	// dbUsr := "dbadmin"
	// dbPass := "AllDone07"
	dbName := "jrsbat"
	// // toRemove := "580c26f78600bf11546eafa0"
	// dbPath := fmt.Sprintf("mongodb://%s:%s@ds061246.mlab.com:61246/%s", dbUsr, dbPass, dbName)
	// session, err := mgo.Dial(dbPath)
	// if err != nil {
	// 	panic(err)
	// }
	// defer session.Close()
	var mc MongoConnection
	err := mc.createConnection()
	if err != nil {
		panic(err)
	}
	sess, col, err := mc.getSessionAndCollection(dbName, "users")
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	mongo := sess.DB(dbName)
	// mongo := mc.session.DB(dbName)

	// Create the value.
	personName := PersonName{
		First: "Chris",
		Last:  "Schweizer",
	}
	person := Person{Name: personName}
	person.mongoID = bson.NewObjectId()
	fmt.Printf("%s\n", person.mongoID.String())
	person.ID = person.mongoID.String()
	// usersCol := mongo.C("users")
	err = col.Insert(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("inserted %s %s with ID %s\n", person.Name.First, person.Name.Last, person.ID)

	// err = usersCol.RemoveId(bson.ObjectIdHex(toRemove))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("removed %s", toRemove)

	fmt.Println(col.Name)
	fmt.Println(personName.First)
	colList, err := mongo.CollectionNames()
	if err != nil {
		panic(err)
	}
	cols := strings.Join(colList, ",")
	fmt.Println(cols)

}

// NOTE Firebase with this client requires us to pass the URL, and token around to the collections, then the collections need t

// PersonName testing
type PersonName struct {
	First string `bson:"First" json:"first"`
	Last  string `bson:"Last" json:"last"`
}

// Person testing
type Person struct {
	mongoID bson.ObjectId
	ID      string     `bson:"ID" json:"id"`
	Name    PersonName `bson:"Name" json:"name"`
}

// MongoConnection used to copy sessions
type MongoConnection struct {
	session *mgo.Session
}

func (c *MongoConnection) createConnection() (err error) {
	fmt.Println("Connecting to cloud mongo server....")
	dbUsr := "dbadmin"
	dbPass := "AllDone07"
	dbName := "jrsbat"
	dbHost := "ds061246.mlab.com"
	dbPort := 61246
	dbPath := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", dbUsr, dbPass, dbHost, dbPort, dbName)
	c.session, err = mgo.Dial(dbPath)
	if err != nil {
		err = fmt.Errorf("Could not create session to %s", dbHost)
	}
	//This will create a unique index to ensure that there won't be duplicate shorturls in the database.
	// index := mgo.Index{
	// 	Key:      []string{"$text:shorturl"},
	// 	Unique:   true,
	// 	DropDups: true,
	// }
	// urlcollection.EnsureIndex(index)
	// }
	return
}

func (c *MongoConnection) getSessionAndCollection(dbStr, col string) (session *mgo.Session, itemCol *mgo.Collection, err error) {
	if c.session != nil {
		session = c.session.Copy()
		itemCol = session.DB(dbStr).C(col)
	} else {
		err = errors.New("No original session found")
	}
	return
}
