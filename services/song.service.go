package services

import (
	"context"
	"encoding/json"
	"test1-tribal/database"
	"test1-tribal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type filter struct {
	name   string
	artist string
	album  string
}

var songCollection *mongo.Collection = database.OpenCollection(database.Client, "song")

func GetSong(name string, artist string, album string) (song *models.Song, err error) {

	//var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)

	songFilter := make(map[string]interface{})

	if len(name) > 0 {
		songFilter["name"] = name
	}

	if len(artist) > 0 {
		songFilter["artist"] = artist
	}

	if len(album) > 0 {
		songFilter["album"] = album
	}
	jsonStr, err := json.Marshal(songFilter)

	var mapData map[string]interface{}
	json.Unmarshal(jsonStr, &mapData)
	err = songCollection.FindOne(context.TODO(), bson.M(mapData)).Decode(&song)

	return
}
