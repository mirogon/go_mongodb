package go_mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbCollectionImpl struct {
	collection *mongo.Collection
}

func CreateMongoDbCollection(mongoCollection *mongo.Collection) MongoDbCollectionImpl {
	return MongoDbCollectionImpl{collection: mongoCollection}
}

func (mongoCollection MongoDbCollectionImpl) InsertOne(data interface{}) error {
	_, err := mongoCollection.collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (mongoCollection MongoDbCollectionImpl) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return mongoCollection.collection.FindOne(context.TODO(), filter, opts[0])
}

func (mongoCollection MongoDbCollectionImpl) Find(filter interface{}, opt *options.FindOptions) (*mongo.Cursor, error) {
	return mongoCollection.collection.Find(context.TODO(), filter, opt)
}

func (mongoCollection MongoDbCollectionImpl) FindAll(filter interface{}) (*mongo.Cursor, error) {
	return mongoCollection.collection.Find(context.TODO(), filter)
}

func (mongoCollection MongoDbCollectionImpl) ReplaceOne(filter interface{}, replacement interface{}) error {
	updateResult, err := mongoCollection.collection.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 0 {
		return errors.New("Found no matches")
	}
	return nil
}

func (mongoCollection MongoDbCollectionImpl) UpsertOne(filter interface{}, replacement interface{}) error {
	upsert := true
	options := options.ReplaceOptions{Upsert: &upsert}
	_, err := mongoCollection.collection.ReplaceOne(context.TODO(), filter, replacement, &options)
	if err != nil {
		return err
	}
	return nil
}
func (mongoCollection MongoDbCollectionImpl) DeleteOne(filter interface{}) error {
	deleteResult, err := mongoCollection.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return err
	}
	return nil
}

func (mongoCollection MongoDbCollectionImpl) DeleteAll(filter interface{}) error {
	deleteResult, err := mongoCollection.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return nil
	}
	return nil
}

func (mongoCollection MongoDbCollectionImpl) Count(filter interface{}) (int64, error) {
	return mongoCollection.collection.CountDocuments(context.TODO(), filter)
}

func (mongoCollection MongoDbCollectionImpl) CreateTTLIndex(ttlSeconds int32) {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"createdAt": 1},
		Options: options.Index().SetExpireAfterSeconds(ttlSeconds),
	}
	mongoCollection.collection.Indexes().CreateOne(context.TODO(), indexModel)
}
