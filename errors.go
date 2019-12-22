package fdb2

import "errors"

var (
	ErrEmptyBasePath       = errors.New("empty base path")
	ErrEmptyItemName       = errors.New("empty item name(check illegal char stripping)")
	ErrEmptyCollectionPath = errors.New("empty collection path(try running collection.init())")
	ErrEmptyDocumentPath   = errors.New("empty path")
	ErrWriteFileTimeout    = errors.New("write file timeout reached")
	ErrDeleteFileTimeout   = errors.New("delete file timeout reached")
)
