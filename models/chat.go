package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	UserId  primitive.ObjectID `bson:"user_id" json:"user_id"`
	Source  string             `bson:"source" json:"source"`
	Type    string             `bson:"type" json:"type"`
	Message string             `bson:"message" json:"message"`
	Time    time.Time          `bson:"time"`
}
