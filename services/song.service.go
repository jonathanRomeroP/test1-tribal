package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"test1-tribal/database"
	"test1-tribal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type filter struct {
	name   string
	artist string
	album  string
}

var songCollection *mongo.Collection = database.OpenCollection(database.Client, "song")

func GetSong(name string, artist string, album string) (songs []*models.SongAny, err error) {

	var ctx, _ = context.WithTimeout(context.TODO(), 100*time.Second)

	filter := bson.M{
		"$or": []bson.M{
			{
				"Name": bson.M{
					"$regex": primitive.Regex{
						Pattern: name,
						Options: "i",
					},
				},
				"Artist": bson.M{
					"$regex": primitive.Regex{
						Pattern: artist,
						Options: "i",
					},
				},
				"Album": bson.M{
					"$regex": primitive.Regex{
						Pattern: album,
						Options: "i",
					},
				},
			},
		},
	}

	cursor, err := songCollection.Find(context.TODO(), filter)

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var song *models.SongAny
		cursor.Decode(&song)
		songs = append(songs, song)
	}

	if songs == nil {
		err = errors.New("No hay datos")
	}

	return
}

func SaveSong(song []any) {
	for _, v := range song {
		jsonbody, _ := json.Marshal(v)
		student := models.Song{}
		json.Unmarshal(jsonbody, &student)
		idSong := student.IdSong
		var res models.Song

		errFind := songCollection.FindOne(context.TODO(), bson.D{{"IdSong", idSong}}).Decode(&res)

		if errFind != nil {
			if idSong, _ := strconv.Atoi(student.IdSong); idSong != 0 {
				songCollection.InsertOne(context.TODO(), v)
				fmt.Printf("este es el id %v\n", idSong)
				fmt.Printf("el id  %v se guardo\n", idSong)
			}
			fmt.Println("el id es 0, por eso no se guardo")
			continue
		}

		fmt.Printf("el id  %v ya existe\n", idSong)
	}

}
