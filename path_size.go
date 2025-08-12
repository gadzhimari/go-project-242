package code

import (
	"errors"
	"fmt"
	"os"
)

func GetSize(path string) (string, error) {
	if path == "" {
		return "", errors.New("path cannot be empty")
	}

	fileInfo, err := os.Lstat(path)

	if err != nil {
		return "", errors.New("failed to get file info")
	}

	var totalSize int64
	if fileInfo.IsDir() {
		files, err := os.ReadDir(path)

		if err != nil {
			return "", errors.New("failed to read directory")
		}

		for _, file := range files {
			if !file.IsDir() {
				fileInfo, err := file.Info()

				if err != nil {
					return "", errors.New("failed to get file info")
				}

				totalSize += fileInfo.Size()
			}
		}
	} else {
		totalSize = fileInfo.Size()
	}

	return fmt.Sprintf("%vB \t%s", totalSize, path), nil
}
