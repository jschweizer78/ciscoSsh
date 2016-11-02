package main

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/jschweizer78/ciscoSsh/interfaces"
)

type userStorage struct {
	name  string
	dbStr string
	ms    *mgo.Session
}

func newUserStorage(db *mgo.Session, dbStr string) (us *userStorage) {
	us = &userStorage{name: "users", ms: db, dbStr: dbStr}
	return
}

func (us *userStorage) getColAndSession() (*mgo.Collection, *mgo.Session) {
	sess := us.ms.Copy()
	col := sess.DB(us.dbStr).C(us.MyName())
	return col, sess
}

func (us *userStorage) AddOne(item interfaces.StorageItem) string {
	usr, found := item.(User)
	if !found {
		panic(fmt.Errorf("Can't assert item %s", item.GetID()))
	}
	usr.HexID = bson.NewObjectId()
	usr.ID = usr.HexID.Hex()

	col, sess := us.getColAndSession()
	defer sess.Close()
	err := col.Insert(usr)
	if err != nil {
		panic(err)
	}

	return usr.ID
}

func (us *userStorage) GetOne(id string) (interfaces.StorageItem, error) {
	var err error
	isHex := bson.IsObjectIdHex(id)
	if isHex {
		err = fmt.Errorf("ID %s is not HEX", id)
		return nil, err
	}
	bID := bson.ObjectIdHex(id)
	var usr User
	col, sess := us.getColAndSession()
	defer sess.Close()
	err = col.FindId(bID).One(usr)
	if err != nil {
		err = fmt.Errorf("Could not find userID %s", id)
		return nil, err
	}
	return usr, nil
}

func (us *userStorage) GetAll() (items []interfaces.StorageItem) {

	return
}

func (us *userStorage) DeleteOne(id string) (err error) {

	return
}

func (us *userStorage) UpdateOne(c interfaces.StorageItem) (err error) {

	return
}

func (us *userStorage) MyName() string {

	return us.name
}

/*

AddOne(c interfaces.StorageItem) string
GetOne(id string) (StorageItem, error)
GetAll() (items []interfaces.StorageItem)
DeleteOne(id string) (err error)
UpdateOne(c interfaces.StorageItem) (err error)
MyName() (id string)

*/

// MongoConnection used to copy sessions
type MongoConnection struct {
	session *mgo.Session
	dbUsr   string
	dbName  string
	dbHost  string
	dbPort  int
}

func (c *MongoConnection) createConnection() (err error) {
	fmt.Println("Connecting to cloud mongo server....")
	c.dbUsr = "dbadmin"
	dbPass := "AllDone07"
	c.dbName = "jrsbat"
	c.dbHost = "ds061246.mlab.com"
	c.dbPort = 61246
	dbPath := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", c.dbUsr, dbPass, c.dbHost, c.dbPort, c.dbName)
	c.session, err = mgo.Dial(dbPath)
	if err != nil {
		err = fmt.Errorf("Could not create session to %s", c.dbHost)
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

func (c *MongoConnection) getSession() (session *mgo.Session, err error) {
	if c.session != nil {
		session = c.session.Copy()
	} else {
		err = c.createConnection()
		if err != nil {
			return
		}
		session = c.session.Copy()
	}
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
