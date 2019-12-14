package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectToDB : connects to mongo database by connection string
func ConnectToDB(conn string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + conn))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// GetAccessToDB : gets access to database
func GetAccessToDB(client *mongo.Client, dbString string) *mongo.Database {
	return client.Database(dbString)
}

// GetAccessToCollection : gets access to collection
func GetAccessToCollection(db *mongo.Database, collString string) *mongo.Collection {
	return db.Collection(collString)
}
