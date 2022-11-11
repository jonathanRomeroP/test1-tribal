package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Name          *string            `json:"name" bson:"name"`
	Email         *string            `json:"email" bson:"email" binding:"required"`
	Password      *string            `json:"password" bson:"password"`
	Token         *string            `json:"token" bson:"token"`
	Refresh_token *string            `json:"refresh_token" bson:"refresh_token" `
}
