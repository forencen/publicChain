package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

type Database interface {
	Close()
	Put(key []byte, value []byte) error
	Delete(key []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	NewBatch() Batch
}

type BatchInterface interface {
	Put(key []byte, value []byte)
	Delete(key []byte)
	Commit() error
	Rollback()
}

type Batch struct {
	db    *leveldb.DB
	batch *leveldb.Batch
}

func (b *Batch) Put(key []byte, value []byte) {
	b.batch.Put(key, value)
}

func (b *Batch) Delete(key []byte) {
	b.batch.Delete(key)
}

func (b *Batch) Commit() error {
	return b.db.Write(b.batch, nil)
}

func (b *Batch) Rollback() {
	b.batch.Reset()
}

type DbHelper struct {
	db *leveldb.DB
}

func NewDbHelper(path string) *DbHelper {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Panic(err)
	}
	return &DbHelper{
		db: db,
	}
}

func (d *DbHelper) Close() {
	d.db.Close()
}

func (d *DbHelper) Put(key []byte, value []byte) error {
	return d.db.Put(key, value, nil)
}

func (d *DbHelper) Delete(key []byte) error {
	return d.db.Delete(key, nil)
}

func (d *DbHelper) Get(key []byte) ([]byte, error) {
	return d.db.Get(key, nil)
}

func (d *DbHelper) Has(key []byte) (bool, error) {
	return d.db.Has(key, nil)
}

func (d *DbHelper) NewBatch() Batch {
	return Batch{
		db:    d.db,
		batch: new(leveldb.Batch),
	}
}
