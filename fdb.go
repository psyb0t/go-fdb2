package fdb2

import (
	"os"
	"path"
)

type FDB struct {
	path string
}

func NewFDB(dbPath string) (*FDB, error) {
	err := os.MkdirAll(dbPath, 0700)
	if err != nil {
		return nil, err
	}

	db := &FDB{path: dbPath}

	return db, nil
}

func (fdb *FDB) NewCollection(name string) (*Collection, error) {
	collectionPath := path.Join(fdb.path, name)
	err := os.MkdirAll(collectionPath, 0700)
	if err != nil {
		return nil, err
	}

	collection := &Collection{
		db:   fdb,
		path: collectionPath,
		name: name,
	}

	return collection, nil
}
