package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	log.Println("aaa")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
