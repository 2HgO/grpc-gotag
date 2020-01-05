package db

import (
	"log"

	codec "github.com/amsokol/mongo-go-driver-protobuf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User ...
var User *Collection

func newCollection(col *mongo.Collection) *Collection {
	return &Collection{
		collection: col,
		database: col.Database(),
	}
}

func init() {
	reg := codec.Register(bson.NewRegistryBuilder()).Build()
	
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/experiment"),
		options.Client().SetRegistry(reg))

	if err != nil {
		log.Fatalln(err)
	}

	client.Connect(NoContext)
	db := client.Database("experiment")
	User = newCollection(db.Collection("test"))
}
