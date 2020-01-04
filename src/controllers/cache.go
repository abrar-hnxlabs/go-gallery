package controllers

import (
	"fmt"
	"path/filepath"
	"os"
)

func EnsureCacheDir(extraDir string) (string, error) {
	path, _ := filepath.Abs("./usercache")
	if len(extraDir) > 0 {
		path = filepath.Join(path, extraDir)
	}
	
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return "" , err
		}
	}
	fmt.Printf("Making cache dir, %s\n", path)
	return path, nil
}