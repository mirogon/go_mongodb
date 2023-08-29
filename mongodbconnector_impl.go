package go_mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbConnectorImpl struct {
}

func (connector MongoDbConnectorImpl) Connect(options *options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		return client, err
	}
	return client, nil
}
