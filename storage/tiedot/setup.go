package tiedot

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/HouzuoGuo/tiedot/db"
)

func doesColExsit(name string, myDB *db.DB) bool {
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

// NewStorageDB used to set loaction and get PTR
func NewStorageDB(path string) *db.DB {
	// (Create if not exist) open a database
	myDB, err := db.OpenDB(path)
	if err != nil {
		panic(err)
	}
	return myDB
}

func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func fillStruct(s interface{}, m map[string]interface{}) error {
	for k, v := range m {
		err := setField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertStringtoID(id string) int {
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(idInt)
}
