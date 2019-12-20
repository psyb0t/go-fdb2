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

// ReadDirFilesWithEndingName list all file names in the given dirpath that end with the given expectedStrEnd
func ReadDirFilesWithEndingName(dirpath, expectedStrEnd string) ([]string, error) {
	resultFiles := []string{}

	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return []string{}, err
	}

	for _, f := range files {
		fname := f.Name()
		if fname[len(fname)-5:] == expectedStrEnd {
			resultFiles = append(resultFiles, fname[:len(fname)-5])
		}
	}

	return resultFiles, nil
}
