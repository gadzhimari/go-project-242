package code

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	ErrPath    = "no such file or directory"
	ErrReadDir = "failed to read directory"
)

func GetPathSize(path string, human, all bool) (string, error) {
	size, err := GetSize(path, all)

	if err != nil {
		return "", err
	}

	formattedSize := formatSize(size, human)

	return formattedSize, nil
}

func GetSize(path string, all bool) (int64, error) {
	var totalSize int64 = 0
	fileInfo, err := os.Lstat(path)
	// fmt.Println("fileInfo ", fileInfo)

	if err != nil {
		return totalSize, errors.New(ErrPath)
	}

	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)

		if err != nil {
			return totalSize, errors.New(ErrReadDir)
		}

		for _, entry := range entries {
			if !all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}

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

func formatSize(size int64, human bool) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIndex := 0
	floatSize := float64(size)

	if !human {
		return fmt.Sprintf("%d%s", size, units[0])
	}

	for floatSize >= 1024 && unitIndex < len(units)-1 {
		floatSize = floatSize / 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%.0f%s", floatSize, units[unitIndex])
	}

	return fmt.Sprintf("%.1f%s", floatSize, units[unitIndex])
}
