package main

import "github.com/gin-gonic/gin"

func router() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/save", saveAlbums)

	router.POST("/albums", postAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		logger.Fatalf(err.Error())
	}
}
