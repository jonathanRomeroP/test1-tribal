package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	IdSong   string             `json:"idSong" bson:"idSong"`
	Name     string             `json:"name" bson:"name"`
	Duration string             `json:"duration" bson:"duration"`
	Album    string             `json:"album" bson:"album"`
	Artwork  string             `json:"artwork" bson:"artwork"`
	Price    string             `json:"price" bson:"price"`
	Origin   string             `json:"origin" bson:"origin"`
}

type SongResponse struct {
	IdSong   string `json:"idSong" bson:"idSong"`
	Name     string `json:"name" bson:"name"`
	Duration string `json:"duration" bson:"duration"`
	Album    string `json:"album" bson:"album"`
	Artwork  string `json:"artwork" bson:"artwork"`
	Price    string `json:"price" bson:"price"`
	Origin   string `json:"origin" bson:"origin"`
}
