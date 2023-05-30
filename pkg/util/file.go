// Package util provides utility functions for working with files and directories.
package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ReadFileP is a convenience function that calls ReadFile and panics if an error occurs.
func ReadFileP(filepath string) string {
	b, err := os.ReadFile(filepath)
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
			println(os.Getwd())
			println(fmt.Sprintf("Warning: ListFilesInDir, error walking path %q: %v", path, err))
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
