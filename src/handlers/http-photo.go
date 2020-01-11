package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"bufio"
)

func GetPhoto(c *gin.Context) {
	filePath := c.Query("f")
	if _, err := os.Stat(filePath); os.IsNotExist(err){
		c.JSON(http.StatusNotFound, gin.H{
			"f": filePath,
		})
		return
	}
	f, _ := os.Open(filePath)
	reader :=  bufio.NewReader(f)
	size, _ := f.Stat()
	c.DataFromReader(http.StatusOK, size.Size(), "application/octet-stream", reader, make(map[string]string))
}