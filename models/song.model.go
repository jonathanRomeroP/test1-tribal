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
	Artist   string             `json:"artist" bson:"artist"`
}

type SongStruct struct {
	IdSong   string `json:"idSong" bson:"idSong"`
	Name     string `json:"name" bson:"name"`
	Duration string `json:"duration" bson:"duration"`
	Album    string `json:"album" bson:"album"`
	Artwork  string `json:"artwork" bson:"artwork"`
	Price    string `json:"price" bson:"price"`
	Origin   string `json:"origin" bson:"origin"`
	Artist   string `json:"artist" bson:"artist"`
}

type SongAny struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	IdSong   any                `json:"idSong" bson:"IdSong"`
	Name     any                `json:"name" bson:"Name"`
	Duration any                `json:"duration" bson:"Duration"`
	Album    any                `json:"album" bson:"Album"`
	Artwork  any                `json:"artwork" bson:"Artwork"`
	Price    any                `json:"price" bson:"Price"`
	Origin   any                `json:"origin" bson:"Origin"`
	Artist   any                `json:"artist" bson:"Artist"`
}
