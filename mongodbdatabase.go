package go_mongodb

type MongoDbDatabase interface {
	GetCollection(string) MongoDbCollection
}
