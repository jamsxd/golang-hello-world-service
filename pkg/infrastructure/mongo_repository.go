package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/jamsxd/golang-hello-world-service/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type basicMongoRepository struct {
	client *mongo.Client
}

func NewBasicMongoRepository() domain.HelloWorldRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Printf("Error trying to connect to mongo: %v", err.Error())
		return nil
	}
	return &basicMongoRepository{
		client: client,
	}
}

func (b *basicMongoRepository) Save(gretting string) error {
	collection := b.client.Database("test").Collection("grettings")
	_, err := collection.InsertOne(context.Background(), bson.D{{"gretting", gretting}})
	if err != nil {
		fmt.Printf("Error trying to save a gret: %v", err.Error())
		return err
	}
	return nil

}
