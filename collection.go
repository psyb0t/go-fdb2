package fdb2

import (
	"os"
	"path"
)

type Collection struct {
	db     *FDB
	parent *Collection
	path   string
	name   string
}

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

func (c *Collection) Delete() error {
	return os.RemoveAll(c.path)
}
