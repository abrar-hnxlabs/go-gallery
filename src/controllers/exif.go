package controllers

import (
	"github.com/rwcarlsen/goexif/exif"
	"os"
	"time"
)
func ExifTime(path string) (time.Time, error){
	f, err := os.Open(path)
	if err != nil {
		return time.Now(), err
	}

	x, err := exif.Decode(f)
	if err != nil {
		return time.Now(), err
	}

	return x.DateTime()
}