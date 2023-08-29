package go_mongodb

type MongoDbClient interface {
	GetDatabase(name string) MongoDbDatabase
}
