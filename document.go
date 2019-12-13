package fdb2

import (
	"io/ioutil"
	"time"
)

type Document struct {
	db         *FDB
	collection *Collection
	path       string
	name       string
	value      []byte
}

func (d *Document) Set(value []byte) error {
	err := WriteFile(d.path, value, time.Second*5)
	if err != nil {
		return err
	}

	return nil
}

func (d *Document) SetString(value string) error {
	return d.Set([]byte(value))
}

func (d *Document) Get() ([]byte, error) {
	value, err := ioutil.ReadFile(d.path)
	if err != nil {
		return value, err
	}

	return value, nil
}

func (d *Document) GetString() (string, error) {
	value, err := d.Get()
	if err != nil {
		return "", err
	}

	return string(value), nil
}

func (d *Document) Delete() error {
	return DeleteFile(d.path, time.Second*5)
}
