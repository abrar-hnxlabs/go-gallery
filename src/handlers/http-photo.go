package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func GetPhoto(c *gin.Context) {
	filePath := c.Query("f")
	if _, err := os.Stat(filePath); os.IsNotExist(err){
		c.JSON(http.StatusNotFound, gin.H{
			"f": filePath,
		})
		return
	}
}