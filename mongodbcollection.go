package go_mongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbCollection interface {
	InsertOne(interface{}) error
	FindOne(filter interface{}) *mongo.SingleResult
	Find(filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error)
	FindAll(filter interface{}) (*mongo.Cursor, error)
	ReplaceOne(filter interface{}, replacement interface{}) error
	UpsertOne(filter interface{}, replacement interface{}) error
	DeleteOne(filter interface{}) error
	DeleteAll(filter interface{}) error
	Count(filter interface{}) (int64, error)
}

func PersistOne(collection MongoDbCollection, data interface{}) error {
	if collection == nil {
		return errors.New("Collection missing")
	}
	err := collection.InsertOne(data)
	if err != nil {
		return err
	}
	return nil
}

// Only replaces, if it doesnt exist already, does nothing and returns error
func ReplaceOne[filterValueType any, replaceType any](collection MongoDbCollection, filterKey string, filterValue filterValueType, newValue replaceType) error {
	if collection == nil {
		return errors.New("Collection missing")
	}
	err := collection.ReplaceOne(bson.D{{filterKey, filterValue}}, newValue)
	if err != nil {
		return err
	}
	return nil
}

func GetOne[filterValueType any, resultType any](collection MongoDbCollection, filterKey string, filterValue filterValueType) (resultType, error) {
	var empty resultType
	if collection == nil {
		return empty, errors.New("Collection missing")
	}
	filter := bson.D{{filterKey, filterValue}}
	result := collection.FindOne(filter)
	if result == nil || result.Err() != nil {
		return empty, errors.New("Not found")
	}
	var data resultType
	err := result.Decode(&data)
	if err != nil {
		return empty, err
	}
	return data, nil
}
func DeleteOne[filterValueType any](collection MongoDbCollection, filterKey string, filterValue filterValueType) error {
	if collection == nil {
		return errors.New("Collection missing")
	}
	filter := bson.D{{filterKey, filterValue}}
	err := collection.DeleteOne(filter)
	if err != nil {
		return err
	}
	return nil
}
func DeleteAll(collection MongoDbCollection) error {
	if collection == nil {
		return errors.New("Collection missing")
	}
	filter := bson.D{}
	err := collection.DeleteAll(filter)
	if err != nil {
		return err
	}
	return nil
}

func Exists[filterValueType any](collection MongoDbCollection, filterKey string, filterValue filterValueType) bool {
	filter := bson.D{{filterKey, filterValue}}
	result := collection.FindOne(filter)
	if result == nil {
		return false
	}
	return result.Err() == nil
}
