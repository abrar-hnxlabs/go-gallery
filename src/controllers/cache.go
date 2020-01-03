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
		fmt.Println(path)
		err := os.MkdirAll(path, os.ModeDir)
		return "", err
	}
	 return path, nil
}