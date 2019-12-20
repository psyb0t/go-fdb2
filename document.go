package fdb2

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

// Document is the object representing a document(key) which will contain a value
type Document struct {
	db         *FDB
	collection *Collection
	basePath   string
	path       string
	name       string
	value      []byte
}

func (d *Document) init() error {
	if d.basePath == "" {
		return ErrEmptyBasePath
	}

	d.name = CleanPathString(d.name)
	if d.name == "" {
		return ErrEmptyItemName
	}

	d.path = path.Join(d.basePath, fmt.Sprintf("%s.%s", d.name, documentExt))

	return nil
}

// Set sets the []byte value of the document
func (d *Document) Set(value []byte) error {
	err := WriteFile(d.path, value, time.Second*5)
	if err != nil {
		return err
	}

	return nil
}

// SetString sets the string value of the document
func (d *Document) SetString(value string) error {
	return d.Set([]byte(value))
}

// Get returns the []byte value of the document
func (d *Document) Get() ([]byte, error) {
	value, err := ioutil.ReadFile(d.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []byte(""), nil
		}

		return []byte(""), err
	}

	return value, nil
}

// GetString returns the string value of the document
func (d *Document) GetString() (string, error) {
	value, err := d.Get()
	if err != nil {
		return "", err
	}

	return string(value), nil
}

// Delete removes the document from the collection
func (d *Document) Delete() error {
	return DeleteFile(d.path, time.Second*5)
}
