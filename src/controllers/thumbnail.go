package controllers

import (
	"github.com/disintegration/imaging"
	"crypto/md5"
	"path/filepath"
	"encoding/hex"
	"io"
	"os"
)

func Thumbnail(path string, size int) (string, error) {
	cacheDir, err := EnsureCacheDir("thumbnails")
	if err != nil {
		return "", err
	}
	
	h := md5.New()
	io.WriteString(h, path)
	savePath := filepath.Join(cacheDir,hex.EncodeToString(h.Sum(nil))+".jpg")

	if _, err := os.Stat(savePath); !os.IsNotExist(err) {
		return savePath, nil
	}

	srcImage, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return "", err
	}

	dstImage := imaging.Resize(srcImage, size, 0, imaging.Lanczos)
	imaging.Save(dstImage, savePath)
	return savePath, nil
}