package main

import (
	"fmt"
	"go-gallery/src/controllers"
	"os"
	// "github.com/gin-gonic/gin"
	// "net/http"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: main {images-dir}")
		os.Exit(1)
	}
	imagesDir := os.Args[1]
	cacheDir, _ := controllers.EnsureCacheDir("")
	scanner, err := controllers.NewScanner(imagesDir, "\\.(jpg|jpeg)$")

	if err != nil {
		fmt.Println(err)
		return
	}
	files := scanner.GetNewFiles()
	store := controllers.NewStore()
	fmt.Printf("Number of files: %d \n", len(files))
	total := len(files)
	for i :=0 ; i < total; i++ {
		thumbnailPath, err:= controllers.Thumbnail(files[i], 250)
		if err != nil {
			continue
		}
		taken, err := controllers.ExifTime(files[i])
		if err == nil {
			store.Add(files[i], thumbnailPath, taken)
		}
		fmt.Printf("Processed file %d / %d \n", i+1, total)
	}

	store.Save(cacheDir+"/store.json")

	// s2 := controllers.NewStore().Load(cacheDir+"/store.json")
	// fmt.Println(s2)
}