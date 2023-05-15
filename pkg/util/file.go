// Package util provides utility functions for working with files and directories.
package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadFile reads the contents of a file given its path and returns the contents as a string.
// If the file cannot be read, the function returns an error.
func ReadFile(filepath string) (string, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("cannot read file %s", filepath)
	}
	return string(b), err
}

// ReadFileP is a convenience function that calls ReadFile and panics if an error occurs.
func ReadFileP(filepath string) string {
	b, err := ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func ListFilesInDir(dir string) ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory %q does not exist", dir)
	}
	var files []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Warning: ListFilesInDir, error walking path %q: %v\n", path, err)
			return nil
		}
		if !info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			relPath, err := filepath.Rel(cwd, absPath)
			if err != nil {
				return err
			}
			relPath = strings.TrimPrefix(relPath, "./")
			files = append(files, filepath.ToSlash(relPath))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// MkdirPAndWriteFile writes the edited content to the file
func MkdirPAndWriteFile(filePath, content string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(filePath, []byte(content), 0644)
}
