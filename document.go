package fdb2

import (
	"io/ioutil"
	"time"
)

// Document is the object representing a document(key) which will contain a value
type Document struct {
	db         *FDB
	collection *Collection
	path       string
	name       string
	value      []byte
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
		return value, err
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
