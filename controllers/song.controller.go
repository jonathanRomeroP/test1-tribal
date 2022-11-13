package controllers

import (
	"fmt"
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
			fmt.Println("Cancion NO esta guardada en local....")
			song := services.GetSongOfAllClients(name, artist, album)
			c.JSON(http.StatusOK, song)
			go services.SaveSong(song)
			return
		}
		fmt.Println("Cancion SI esta guardada en local....")
		c.JSON(http.StatusOK, song)
	}

}
