package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-gallery/src/controllers"
)

func GetData(c *gin.Context) {
	cacheDir, _ := controllers.EnsureCacheDir("")
	store := controllers.NewStore(cacheDir+"/store.json")
	c.JSON(http.StatusOK, store)
}