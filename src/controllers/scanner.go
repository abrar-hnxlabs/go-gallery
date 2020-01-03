package controllers

import (
  "path/filepath"
  "os"
  "regexp"
  "fmt"
)

type scanner struct{
  rootFolderAbs string
  files []string
  fileFilterRegex *regexp.Regexp
}

func NewScanner(root string, searchPattern string) (*scanner, error ){
  rootAbs, _ := filepath.Abs(root)
  stat , err := os.Lstat(rootAbs)
  if err != nil {
    return nil, err
  }

  if !stat.IsDir() {
    return nil, fmt.Errorf("root is not a directory/or is symlink")
  }

  s := scanner{rootFolderAbs: rootAbs}
  s.fileFilterRegex = regexp.MustCompile(searchPattern)
  return &s, nil
}

func (s *scanner) GetNewFiles() ([]string) {
  s.files = make([]string, 0)
  filepath.Walk(s.rootFolderAbs, s.walkFunc)
  return s.files
}

func (s *scanner) walkFunc(path string, info os.FileInfo, err error) error {
  
  if s.fileFilterRegex.MatchString(path) {
    s.files = append(s.files, path)
  }
  return err
}
