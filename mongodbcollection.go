package go_mongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbCollection interface {
	InsertOne(interface{}) error
	FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(filter interface{}, opts *options.FindOptions) (*mongo.Cursor, error)
	FindAll(filter interface{}) (*mongo.Cursor, error)
	ReplaceOne(filter interface{}, replacement interface{}) error
	UpsertOne(filter interface{}, replacement interface{}) error
	DeleteOne(filter interface{}) error
	DeleteAll(filter interface{}) error
	Count(filter interface{}) (int64, error)
	Distinct(fieldName string) ([]interface{}, error)
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
		return errors.New("collection missing")
	}
	err := collection.ReplaceOne(bson.D{{filterKey, filterValue}}, newValue)
	if err != nil {
		return err
	}
	return nil
}

func ReplaceOnce_(collection MongoDbCollection, keyName string, filterValue interface{}, newValue interface{}) error {
	if collection == nil {
		return errors.New("collection missing")
	}

	err := collection.ReplaceOne(bson.D{{keyName, filterValue}}, newValue)
	if err != nil {
		return err
	}

	return nil
}

func GetOne[resultType any](collection MongoDbCollection, filterKey string, filterValue interface{}) (resultType, error) {
	var empty resultType
	if collection == nil {
		return empty, errors.New("collection missing")
	}
	filter := bson.D{{filterKey, filterValue}}
	result := collection.FindOne(filter)
	if result == nil || result.Err() != nil {
		return empty, errors.New("not found")
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

func DeleteOne_(collection MongoDbCollection, keyName string, filterValue interface{}) error {
	if collection == nil {
		return errors.New("collection missing")
	}
	filter := bson.D{{keyName, filterValue}}

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

func Exists(collection MongoDbCollection, keyName string, filterValue interface{}) bool {
	filter := bson.D{{keyName, filterValue}}
	found, err := collection.Count(filter)
	if err != nil {
		return false
	}
	return found > 0
}
