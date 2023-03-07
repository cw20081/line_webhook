package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `bson:"_id" json:"id"`
	Sources map[string]string  `bson:"sources" json:"sources"`
}
