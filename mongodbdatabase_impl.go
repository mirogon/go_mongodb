package go_mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbDatabaseImpl struct {
	database *mongo.Database
}

func CreateMongoDbDatabase(db *mongo.Database) MongoDbDatabaseImpl {
	mongoDbDatabase := MongoDbDatabaseImpl{}
	mongoDbDatabase.database = db
	return mongoDbDatabase
}

func (mongoDatabase MongoDbDatabaseImpl) GetCollection(name string) MongoDbCollection {
	collection := CreateMongoDbCollection(mongoDatabase.database.Collection(name))
	return collection
}
