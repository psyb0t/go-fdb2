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

// Document returns a Document object reference with the given name
func (c *Collection) Document(name string) *Document {
	name = CleanPathString(name)

	documentPath := path.Join(c.path, name)

	document := &Document{
		db:         c.db,
		collection: c,
		path:       documentPath,
		name:       name,
	}

	return document
}

// Collection initializes a collection directory as a child to the parent collection
// and returns a Collection object reference
func (c *Collection) Collection(name string) (*Collection, error) {
	name = CleanPathString(name)

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

// Delete removes the collection and all of its children
func (c *Collection) Delete() error {
	// todo directory locking to stop documents from being added/updated
	return os.RemoveAll(c.path)
}
