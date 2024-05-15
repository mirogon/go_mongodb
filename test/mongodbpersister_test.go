package go_mongodb_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	db_mongo "github.com/mirogon/go_mongodb"
	mock_db "github.com/mirogon/go_mongodb/mocks"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountData struct {
	Id                     uint64 `bson:"_id,omitempty"`
	Email                  string
	Username               string
	UserSince              int64
	Permission             int
	Salt                   string
	PwSaltedHashed         string
	LoginToken             string
	LoginTokenCreationTime int64
	AvailableSessionDate   int64
	Subtoken               int
	Extratoken             int
}

func TestPersistOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)
	accData := AccountData{Email: "test@example.com"}
	mockCollection.EXPECT().InsertOne(accData)

	err := db_mongo.PersistOne(mockCollection, accData)

	if err != nil {
		t.Error()
	}
}

func TestPersistOne_MissingCollection(t *testing.T) {
	accData := AccountData{Email: "test@example.com"}
	err := db_mongo.PersistOne(nil, accData)
	if err == nil {
		t.Error()
	}
}

func TestPersistOne_InsertOneFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)
	accData := AccountData{Email: "test@example.com"}
	mockCollection.EXPECT().InsertOne(accData).Return(errors.New(""))

	err := db_mongo.PersistOne(mockCollection, accData)

	if err == nil {
		t.Error()
	}
}

func TestReplaceOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)
	accData := AccountData{Email: "test@example.com", Id: 10}

	mockCollection.EXPECT().ReplaceOne(bson.D{{"accountid", uint64(10)}}, accData).Return(nil)

	err := db_mongo.ReplaceOne(mockCollection, "accountid", uint64(10), accData)
	if err != nil {
		t.Error()
	}
}

func TestReplaceOne_MissingCollection(t *testing.T) {
	accData := AccountData{Email: "test@example.com", Id: 10}
	err := db_mongo.ReplaceOne(nil, "accountid", uint64(10), accData)
	if err == nil {
		t.Error()
	}
}

func TestReplaceOne_ReplaceOneFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)
	accData := AccountData{Email: "test@example.com", Id: 10}

	mockCollection.EXPECT().ReplaceOne(bson.D{{"accountid", uint64(10)}}, accData).Return(errors.New(""))

	err := db_mongo.ReplaceOne(mockCollection, "accountid", uint64(10), accData)
	if err == nil {
		t.Error()
	}
}

func TestGetOneSpecific(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)

	accData := AccountData{Email: "test@example.com"}
	resultingDoc := mongo.NewSingleResultFromDocument(accData, nil, nil)
	mockCollection.EXPECT().FindOne(bson.D{{"accountid", uint64(55)}}).Return(resultingDoc)

	result, err := db_mongo.GetOne[AccountData](mockCollection, "accountid", 55)
	if err != nil {
		t.Error()
	}
	if result.Email != "test@example.com" {
		t.Error()
	}
}
func TestGetOneSpecific_MissingCollection(t *testing.T) {
	_, err := db_mongo.GetOne[AccountData](nil, "accountid", 55)
	if err == nil {
		t.Error()
	}
}

func TestGetOneSpecific_FindOneFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)

	resultingDoc := mongo.NewSingleResultFromDocument(nil, nil, nil)
	mockCollection.EXPECT().FindOne(bson.D{{"accountid", uint64(55)}}).Return(resultingDoc)

	_, err := db_mongo.GetOne[AccountData](mockCollection, "accountid", 55)
	if err == nil {
		t.Error()
	}
}

func TestDeleteOneSpecific(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)
	mockCollection.EXPECT().DeleteOne(bson.D{{"accountid", 55}}).Return(nil)

	err := db_mongo.DeleteOne(mockCollection, "accountid", 55)
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteOneSpecific_Error(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)
	mockCollection.EXPECT().DeleteOne(bson.D{{"accountid", uint64(5)}}).Return(errors.New(""))

	err := db_mongo.DeleteOne(mockCollection, "accountid", uint64(5))
	if err == nil {
		t.Error(err)
	}
}

func TestExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)
	accData := AccountData{}
	result := mongo.NewSingleResultFromDocument(accData, nil, nil)
	mockCollection.EXPECT().FindOne(bson.D{{"accountid", uint64(10)}}).Return(result)
	exists := db_mongo.Exists(mockCollection, "accountid", uint64(10))
	if !exists {
		t.Error()
	}
}

func TestExists_DoesntExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCollection := mock_db.NewMockMongoDbCollection(ctrl)
	mockCollection.EXPECT().FindOne(bson.D{{"accountid", uint64(10)}}).Return(nil)
	exists := db_mongo.Exists(mockCollection, "accountid", uint64(10))
	if exists {
		t.Error()
	}
}
