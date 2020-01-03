package main

import (
	"fmt"
	"go-gallery/src/controllers"
	// "github.com/gin-gonic/gin"
	// "net/http"
)

func main(){
	cacheDir, _ := controllers.EnsureCacheDir("")
	scanner, err := controllers.NewScanner("./images", "\\.(jpg|jpeg)$")

	if err != nil {
		fmt.Println(err)
		return
	}
	files := scanner.GetNewFiles()
	store := controllers.NewStore()
	for i :=0 ; i <len(files); i++ {
		thumbnailPath, err:= controllers.Thumbnail(files[i], 250)
		if err != nil {
			continue
		}
		taken, err := controllers.ExifTime(files[i])
		if err == nil {
			store.Add(files[i], thumbnailPath, taken)
		}	
	}

	store.Save(cacheDir+"/store.json")

	s2 := controllers.NewStore().Load(cacheDir+"/store.json")
	fmt.Println(s2)
}