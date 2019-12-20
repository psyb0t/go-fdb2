package fdb2

import (
	"fmt"
	"os"
	"path"
)

// Collection is the object representing a collection of collections and/or documents
type Collection struct {
	db       *FDB
	parent   *Collection
	basePath string
	path     string
	name     string
}

func (c *Collection) init() error {
	if c.basePath == "" {
		return ErrEmptyBasePath
	}

	c.name = CleanPathString(c.name)
	if c.name == "" {
		return ErrEmptyItemName
	}

	c.path = path.Join(c.basePath, fmt.Sprintf("%s.%s", c.name, collectionExt))

	return os.MkdirAll(c.path, 0700)
}

// Document returns a Document object reference with the given name
func (c *Collection) Document(name string) (*Document, error) {
	if c.path == "" {
		return nil, ErrEmptyCollectionPath
	}

	document := &Document{
		db:         c.db,
		collection: c,
		basePath:   c.path,
		name:       name,
	}

	err := document.init()
	if err != nil {
		return nil, err
	}

	return document, nil
}

// ListDocuments returns a list of document names created under the collection path
func (c *Collection) ListDocuments() ([]string, error) {
	collectionNames := []string{}

	expectedStrEnd := fmt.Sprintf(".%s", documentExt)
	collectionNames, err := ReadDirFilesWithEndingName(c.path, expectedStrEnd)
	if err != nil {
		return []string{}, err
	}

	return collectionNames, nil
}

// Collection initializes a collection directory as a child to the parent collection
// and returns a Collection object reference
func (c *Collection) Collection(name string) (*Collection, error) {
	if c.path == "" {
		return nil, ErrEmptyCollectionPath
	}

	collection := &Collection{
		db:       c.db,
		parent:   c,
		basePath: c.path,
		name:     name,
	}

	err := collection.init()
	if err != nil {
		return nil, err
	}

	return collection, nil
}

// ListCollections returns a list of collection names created under the collection path
func (c *Collection) ListCollections() ([]string, error) {
	collectionNames := []string{}

	expectedStrEnd := fmt.Sprintf(".%s", collectionExt)
	collectionNames, err := ReadDirFilesWithEndingName(c.path, expectedStrEnd)
	if err != nil {
		return []string{}, err
	}

	return collectionNames, nil
}

// Delete removes the collection and all of its children
func (c *Collection) Delete() error {
	// todo directory locking to stop documents from being added/updated
	return os.RemoveAll(c.path)
}
