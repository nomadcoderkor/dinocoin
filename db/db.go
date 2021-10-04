package db

import (
	"github.com/boltdb/bolt"
	"github.com/nomadcoderkor/dinocoin/utils"
)

// 블록체인과 연계해서 많은 작업을 하게될 Package
const (
	dbName       = "dinochain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

// DB Connection
func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err := tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)

	}
	return db
}
