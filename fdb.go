package fdb2

import (
	"fmt"
	"os"
	"strings"
)

const (
	collectionExt = "fdbc"
	documentExt   = "fdbd"
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

// Collection initializes a collection directory and returns a Collection object reference
func (fdb *FDB) Collection(name string) (*Collection, error) {
	collection := &Collection{
		db:       fdb,
		basePath: fdb.path,
		name:     name,
	}

	err := collection.init()
	if err != nil {
		return nil, err
	}

	return collection, nil
}

// CollectionSequence initializes a collection sequence and returns the last one
func (fdb *FDB) CollectionSequence(sequenceName string) (*Collection, error) {
	var parentCollection *Collection
	var lastCollection *Collection

	collectionNames := strings.Split(sequenceName, "/")
	for _, collectionName := range collectionNames {
		var collection *Collection
		var err error

		if parentCollection == nil {
			collection, err = fdb.Collection(collectionName)
			if err != nil {
				return nil, err
			}

			parentCollection = collection
			lastCollection = collection

			continue
		}

		collection, err = parentCollection.Collection(collectionName)
		if err != nil {
			return nil, err
		}

		parentCollection = collection
		lastCollection = collection
	}

	return lastCollection, nil
}

// ListCollections returns a list of collection names created under the db path
func (fdb *FDB) ListCollections() ([]string, error) {
	collectionNames := []string{}

	expectedStrEnd := fmt.Sprintf(".%s", collectionExt)
	collectionNames, err := ReadDirFilesWithEndingName(fdb.path, expectedStrEnd)
	if err != nil {
		return []string{}, err
	}

	return collectionNames, nil
}
