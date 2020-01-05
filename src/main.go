package main

import (
	"go-gallery/src/controllers"
	"go-gallery/src/handlers"
	"github.com/gin-gonic/gin"
)

func main(){
	thumbsDir, _ := controllers.EnsureCacheDir("thumbnails")
	// start the server
	router := gin.Default()
	
	apiGroup := router.Group("/api")
	apiGroup.Static("/thumbnail", thumbsDir)
	apiGroup.GET("/photo", handlers.GetPhoto)
	apiGroup.GET("/data", handlers.GetData)

	router.Static("/ui", "./src/ui/build")
	go handlers.InitScan()
	router.Run(":3001")
}