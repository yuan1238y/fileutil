package fileutil

import (
	"errors"
	"os"
	"path"
)

// Touch function used for create an empty file
// and create it parent dir
// like: mkdir -p /a/b; touch /a/b/c
func Touch(filepath string, perm os.FileMode) error {
	_, err := os.Stat(filepath)
	if os.IsExist(err) {
		return nil
	}
	// File not exists, check parent directory
	parentDir := path.Dir(filepath)
	fileInfo, err := os.Stat(parentDir)
	if os.IsExist(err) && !fileInfo.IsDir() {
		return errors.New("Parent not a directory")
	} else if os.IsNotExist(err) {
		os.MkdirAll(parentDir, perm)
	}

	// Create empty file
	file, err := os.OpenFile(filepath, os.O_CREATE, perm)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}
