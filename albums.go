package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	ReleaseYear int    `json:"price"`
}

var albums = []album{
	{ID: 1, Title: "A Thousand Suns", Artist: "Linkin Park", ReleaseYear: 2010},
	{ID: 2, Title: "Endless Summer Vacation", Artist: "Miley Cyrus", ReleaseYear: 2023},
	{ID: 3, Title: "The Human Condition", Artist: "Jon Bellion", ReleaseYear: 2016},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum album
	err := ctx.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func saveAlbums(ctx *gin.Context) {
	generateYAMLFile(albums)
	ctx.IndentedJSON(http.StatusOK, albums)
	getYAMLFile()
}
