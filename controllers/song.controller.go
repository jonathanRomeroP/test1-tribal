package controllers

import (
	"net/http"
	"test1-tribal/services"

	"github.com/gin-gonic/gin"
)

func GetAllSong() {

}

func FilterSong() gin.HandlerFunc {

	return func(c *gin.Context) {
		name := c.Query("name")
		artist := c.Query("artist")
		album := c.Query("album")

		song, err := services.GetSong(name, artist, album)
		if err != nil {
			c.JSON(http.StatusOK, services.GetSongOfAllClients(name, artist, album))
			return
		}
		c.JSON(http.StatusOK, song)
	}

}
