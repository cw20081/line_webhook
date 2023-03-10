package repositories

import (
	"context"
	"time"
	"webhook/configs"
	"webhook/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return userCollection.InsertOne(ctx, user)

}

func GetUser(userID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	err := userCollection.FindOne(ctx, bson.D{{"sources.line", userID}}).Decode(&user)

	return user, err
}
