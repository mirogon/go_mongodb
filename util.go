package go_mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoDbHistoryPagingOption(startDatePropertyName string, page int, docsPerPage int) *options.FindOptions {
	opts := options.Find().SetSort(bson.M{startDatePropertyName: -1}).SetSkip(int64(page-1) * int64(docsPerPage)).SetLimit(int64(docsPerPage))
	return opts
}
