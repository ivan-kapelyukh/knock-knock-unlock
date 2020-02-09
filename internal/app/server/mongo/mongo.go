package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func alreadyRegisteredError(username string) error {
	return fmt.Errorf("user %v already registered", username)
}

func New(host string, port int) (*Mongo, error) {
	options := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", host, port))

	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	collection := client.Database("kku").Collection("users")

	return &Mongo{client, collection}, nil
}

func (m *Mongo) Insert(user User) error {
	count, err := m.collection.CountDocuments(context.TODO(), bson.D{{"username", user.Username}})

	if err != nil {
		return err
	}

	if count != 0 {
		return alreadyRegisteredError(user.Username)
	}

	_, err = m.collection.InsertOne(context.TODO(), user)
	return err
}

func (m *Mongo) Find(username string) (*User, error) {
	var user User

	err := m.collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
