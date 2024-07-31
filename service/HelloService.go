package service

import (
	"goweb/service/pojo"
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

// Multipart/Urlencoded Form
func Hello_formdata(c *gin.Context) {
	message := c.PostForm("message")
	name := c.DefaultPostForm("name", "anonymous")

	log.Println("aaa")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello," + name + "_" + message,
	})
}

func Hello_jsondata(c *gin.Context) {
	var req pojo.HelloReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "hello," + req.Account,
	})
}

// 路径参数演示 xxxx.com/:name
func Hello_pathparam(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"message": "hello," + name,
	})
}

// 查询参数演示  xxxx.com?name=aaa
func Hello_queryparam(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")

	c.JSON(http.StatusOK, gin.H{
		"message": "hello," + name,
	})
}
