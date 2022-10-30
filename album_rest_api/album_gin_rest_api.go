package albumRestApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunAlbumGinApi() {
	println("---===== Tutorial: Developing a RESTful API with Go and Gin =====---")

	// 3 lines that do the same:
	// var router *gin.Engine = gin.Default()
	// var router = gin.Default()
	router := gin.Default()

	router.GET("/albums", getAlbums_Gin)
	router.GET("/albums/:id", getAlbumByID_Gin)
	router.POST("/albums", postAlbums_Gin)

	router.Run(serverHostname)
}

func getAlbums_Gin(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums_Gin(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums.AddAlbum(newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID_Gin(c *gin.Context) {
	id := c.Param("id")

	album, ok := albums.GetAlbumByID(id)
	if ok {
		c.IndentedJSON(http.StatusOK, album)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}
