package database

import (
	"errors"

	bolt "go.etcd.io/bbolt"
)

type Database interface {
	Exists(string) bool
	Get(string) string
	Insert(key, value string) error
}

type kvDatabase struct {
	kv *bolt.DB
}

func NewDatabase(path string) (Database, error) {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &kvDatabase{
		kv: db,
	}, nil
}

func (db *kvDatabase) Exists(key string) bool {
	if v := db.Get(key); v != "" {
		return true
	}

	return false
}

func (db *kvDatabase) Get(key string) string {
	var res []byte
	db.kv.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("main"))
		if bucket == nil {
			return errors.New("bucket main not found")
		}

		res = bucket.Get([]byte(key))

		return nil
	})

	return string(res)
}

func (db *kvDatabase) Insert(key, value string) error {
	return db.kv.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("main"))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		return nil
	})
}
