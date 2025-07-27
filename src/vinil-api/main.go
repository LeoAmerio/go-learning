package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Abbey Road", Artist: "The Beatles", Price: 19.99},
	{ID: "2", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 18.99},
	{ID: "3", Title: "Back in Black", Artist: "AC/DC", Price: 17.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func modifyAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			var updatedAlbum Album
			if err := c.BindJSON(&updatedAlbum); err != nil {
				return
			}
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, updatedAlbum)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	// fmt.Println("Welcome to Vinil API")
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", modifyAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")
}
