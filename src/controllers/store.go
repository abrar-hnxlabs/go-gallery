package controllers

import (
	"path/filepath"
	"time"
	"encoding/json"
	"os"
	"io/ioutil"
)


type fileTuple struct {
	Dir string `json:"dir"`
	File string `json:"file"`
	Thumbnail string `json:"thumbnail"`
	Year int `json:"year"`
	Month int `json:"month"`
}

type store struct {
	Data []fileTuple `json:"data"`
}

func NewStore(path string) *store {
	s := store{Data: make([]fileTuple, 0)}
	if len(path) > 0 {
		data, err := ioutil.ReadFile(path)
		if err == nil {
			json.Unmarshal(data, &s)
		}
	}
	return &s
}

func (s *store) Add(path string, thumbnail string, timetaken time.Time){
	f := fileTuple{
		Dir: filepath.Base(filepath.Dir(path)), 
		File: "/api/photo?f="+path, 
		Thumbnail: "/api/thumbnail/"+filepath.Base(thumbnail), 
		Year: timetaken.Year(), 
		Month: int(timetaken.Month())}

	for i :=0; i<len(s.Data); i++ {
		if s.Data[i].File == path {
			s.Data[i] = f
			return
		}
	}
	s.Data = append(s.Data, f)
}

func (s *store) Save(path string) error{
	b, err := json.Marshal(s)
	if err !=nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	f.Write(b)
	f.Close()
	return nil
}