package main

import (
	"LeetCode/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(middleware.Authenticate)

	router.GET("/getData", getData)
	router.GET("/getData1", getData1)
	router.GET("/getData2", getData2)
	// router.GET("/getUrlData/:name/:age", getUrlData)

	router.Run()
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "GET Hiiiiiiiiii",
	})
}

func getData1(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "GET1 Hiiiiiiiiii",
	})
}

func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "GET2 Hiiiiiiiiii",
	})
}
