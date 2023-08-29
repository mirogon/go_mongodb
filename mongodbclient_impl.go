package go_mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoDbClientImpl struct {
	mongoDbClient *mongo.Client
}

func CreateMongoDbClient(mongoClient *mongo.Client) MongoDbClientImpl {
	return MongoDbClientImpl{mongoDbClient: mongoClient}
}

func (client MongoDbClientImpl) GetDatabase(name string) MongoDbDatabase {
	database := CreateMongoDbDatabase(client.mongoDbClient.Database(name))
	return database
}
