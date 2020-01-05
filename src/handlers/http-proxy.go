package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StaticProxy(c *gin.Context){
	getProxyData("/static","/"+c.Param("action"), c)
}

func Proxy(c *gin.Context){
	getProxyData("/","", c)
}

func SockJSProxy(c *gin.Context){
	getProxyData("/sockjs-node","", c)
}

func getProxyData(base string, action string, c *gin.Context){
	response, err := http.Get("http://localhost:3000"+base+action)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, nil)
}