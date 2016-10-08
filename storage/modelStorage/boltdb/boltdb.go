package boltdb

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/jschweizer78/ciscoSsh/models"
)

// BoltStorage to comply with storageInterface for use with boltdb "github.com/boltdb/bolt"
type BoltStorage struct {
	DB     *bolt.DB
	Bucket string
	Prefix string
}

// BoltSwitchStorage to comply with STInterface for use with boltdb "github.com/boltdb/bolt"
type BoltSwitchStorage struct {
	BoltStorage
}

// TODO figure out how to use go func/channels for batch processing

func (bs *BoltStorage) createID(suffix string) string {
	return fmt.Sprintf("%s-%s", bs.Prefix, suffix)
}

// CreateOne to comply with STInterface for use with boltdb
func (bsws *BoltSwitchStorage) CreateOne(sw models.Switch) (string, error) {
	idByte, _ := b.NextSequence()
	id = fmt.Sprintf("%s-%d", bsws.Prefix, int(idByte))
	err := bsws.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bsws.Bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}

		buff, _ := json.Marshal(sw)
		err := b.Put([]byte(id), buff)
		return err
	})

	return id, err
}

// CreateMany to comply with STInterface for use with boltdb
func (bsws *BoltSwitchStorage) CreateMany(sw []models.Switch) ([]string, error) {
	b, err := tx.CreateBucketIfNotExists([]byte(bsws.Bucket))
	if err != nil {
		return fmt.Errorf("create bucket: %v", err)
	}
	var ids []string
	for _, item := range sw {
		idByte, _ := b.NextSequence()
		ids = append(ids, fmt.Sprintf("%s-%d", bsws.Prefix, int(idByte)))
		err = bsws.DB.Update(func(tx *bolt.Tx) error {
			buff, _ := json.Marshal(item)
			err = b.Put([]byte(id), buff)
			return err
		})
	}
	return ids, err
}

// ReadOne to comply with STInterface for use with boltdb
func (bsws *BoltSwitchStorage) ReadOne(id string) (models.Switch, error) {
	var sw models.Switch
	err := bsws.DB.View(func(tx *bolt.Tx) {
		b, err := tx.Bucket([]byte(bsws.Bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		swBuff := b.Get([]byte(id))
		sw = models.Switch(swBuff)
	})
	return sw, err
}

// ReadMany to comply with STInterface for use with boltdb
func (bsws *BoltSwitchStorage) ReadMany(ids []string) ([]models.Switch, error) {

	return nil, nil
}

// UpdateOne to comply with STInterface for use with boltdb
func (bsws *BoltSwitchStorage) UpdateOne(sw models.Switch) error {

	return nil
}

// UpdateMany to comply with STInterface for use with boltdb
func (bsws *BoltSwitchStorage) UpdateMany(sws []models.Switch) error {

}

// DeleteOne to comply with STInterface for use with boltdb
func (bsws *BoltSwitchStorage) DeleteOne(sw string) error {

	return nil
}

// DeleteMany to comply with STInterface for use with boltdb
func (bsws *BoltSwitchStorage) DeleteMany(sws []string) error {

	return nil
}
