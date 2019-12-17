package fdb2

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// CleapPathStrings removes illegal characters from the given path string
func CleanPathString(pathString string) string {
	itemsToClean := []string{"/", "\n", "\r\n", "\\"}
	for _, itemToClean := range itemsToClean {
		pathString = strings.Replace(pathString, itemToClean, "", -1)
	}

	pathString = strings.TrimSpace(pathString)

	return pathString
}

// WriteFile writes to a file locking via a .tmp file
func WriteFile(fpath string, data []byte, timeout time.Duration) error {
	tmpFpath := fmt.Sprintf("%s.tmp", fpath)

	startTime := time.Now()
	for startTime.Add(timeout).After(time.Now()) {
		if _, err := os.Stat(tmpFpath); os.IsNotExist(err) {
			err = ioutil.WriteFile(tmpFpath, data, 0644)
			if err != nil {
				return err
			}

			err = os.Rename(tmpFpath, fpath)
			if err != nil {
				return err
			}

			return nil
		}

		time.Sleep(time.Nanosecond)
	}

	return ErrWriteFileTimeout
}

// DeleteFile removes a non-locked file
func DeleteFile(fpath string, timeout time.Duration) error {
	tmpFpath := fmt.Sprintf("%s.tmp", fpath)

	startTime := time.Now()
	for startTime.Add(timeout).After(time.Now()) {
		if _, err := os.Stat(tmpFpath); os.IsNotExist(err) {
			return os.RemoveAll(fpath)
		}

		time.Sleep(time.Nanosecond)
	}

	return ErrDeleteFileTimeout
}
