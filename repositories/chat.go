package repositories

import (
	"context"
	"log"
	"time"
	"webhook/configs"
	"webhook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var chatCollection *mongo.Collection = configs.GetCollection(configs.DB, "chats")

func SaveChat(chat models.Chat) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return chatCollection.InsertOne(ctx, chat)
}

func IndexChat(user models.User) []models.Chat {
	var res []models.Chat

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := chatCollection.Find(ctx, bson.D{{"user_id", user.Id}, {"type", "recive"}})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &res); err != nil {
		log.Fatal(err)
	}

	return res
}
