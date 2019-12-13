package fdb2

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

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
