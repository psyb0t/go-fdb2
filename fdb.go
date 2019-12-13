package fdb2

import (
	"os"
	"path"
)

// FDB is the file database object which initializes the base db directory and
// is used to create collections
type FDB struct {
	path string
}

// NewFDB creates the db dir and returns a new FDB object reference
func NewFDB(dbPath string) (*FDB, error) {
	err := os.MkdirAll(dbPath, 0700)
	if err != nil {
		return nil, err
	}

	db := &FDB{path: dbPath}

	return db, nil
}

// NewCollection initializes a collection directory and returns a Collection object reference
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
