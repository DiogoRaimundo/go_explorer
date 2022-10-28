package albumRestApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const serverHostname = "localhost:8080"

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func RunAlbumApi() {
	println("---===== Tutorial: Developing a RESTful API with Go and Gin =====---")

	// 3 lines that do the same:
	// var router *gin.Engine = gin.Default()
	// var router = gin.Default()
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run(serverHostname)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// For vs For-each:
	// albumLength := len(albums)
	// for currIdx := 0; currIdx < albumLength; currIdx++ {
	// 	var currAlbum = albums[currIdx]

	// Use _ in place of currIdx to skip the index value
	// To skip the slice item, just omit currAlbum
	for currIdx, currAlbum := range albums {
		if currAlbum.ID == id {
			println("Found at index", currIdx)

			c.IndentedJSON(http.StatusOK, currAlbum)

			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}