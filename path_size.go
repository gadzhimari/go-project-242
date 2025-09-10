package code

import (
	"errors"
	"fmt"
	"os"
)

const (
	ErrPath    = "no such file or directory"
	ErrReadDir = "failed to read directory"
)

func GetSize(path string) (int64, error) {
	var totalSize int64 = 0
	fileInfo, err := os.Lstat(path)

	if err != nil {
		return totalSize, errors.New(ErrPath)
	}

	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)

		if err != nil {
			return totalSize, errors.New(ErrReadDir)
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				entryInfo, err := entry.Info()

				if err != nil {
					return totalSize, errors.New(ErrPath)
				}

				totalSize += entryInfo.Size()
			}
		}
	} else {
		totalSize = fileInfo.Size()
	}

	return totalSize, nil
}

func FormatSize(path string) (string, error) {
	size, err := GetSize(path)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%vB \t%s", size, path), nil
}
