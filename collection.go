package fdb2

import (
	"os"
	"path"
)

// Collection is the object representing a collection of collections and/or documents
type Collection struct {
	db     *FDB
	parent *Collection
	path   string
	name   string
}

// NewDocument returns a new Document object reference with the given name
func (c *Collection) NewDocument(name string) *Document {
	documentPath := path.Join(c.path, name)

	document := &Document{
		db:         c.db,
		collection: c,
		path:       documentPath,
		name:       name,
	}

	return document
}

// NewCollection initializes a collection directory as a child to the parent collection
// and returns a Collection object reference
func (c *Collection) NewCollection(name string) (*Collection, error) {
	collectionPath := path.Join(c.path, name)
	err := os.MkdirAll(collectionPath, 0700)
	if err != nil {
		return nil, err
	}

	collection := &Collection{
		db:     c.db,
		parent: c,
		path:   collectionPath,
		name:   name,
	}

	return collection, nil
}

// Delete removes the collection and all its children
func (c *Collection) Delete() error {
	return os.RemoveAll(c.path)
}
