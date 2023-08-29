package go_mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbConnector interface {
	Connect(*options.ClientOptions) (*mongo.Client, error)
}
