package fdb2

import "errors"

var (
	ErrWriteFileTimeout  = errors.New("write file timeout reached")
	ErrDeleteFileTimeout = errors.New("delete file timeout reached")
)
