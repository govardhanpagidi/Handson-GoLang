package main

import (
	"context"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main(){
 TestTransactionCommit()

}

func TestTransactionCommit() {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()
	var id = primitive.NewObjectID()
	var doc = bson.M{"_id": id, "hometown": "Atlanta", "year": int32(1998)}
	var result *mongo.UpdateResult
	var session mongo.Session
	var update = bson.D{{Key: "$set", Value: bson.D{{Key: "year", Value: int32(2000)}}}}
	if client, err = getMongoClient(); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection = client.Database(dbName).Collection(collectionExamples)
	if _, err = collection.InsertOne(ctx, doc); err != nil {
		log.Fatal(err)
	}

	if session, err = client.StartSession(); err != nil {
		log.Fatal(err)
	}
	if err = session.StartTransaction(); err != nil {
		log.Fatal(err)
	}
	if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if result, err = collection.UpdateOne(sc, bson.M{"_id": id}, update); err != nil {
			log.Fatal(err)
		}
		if result.MatchedCount != 1 || result.ModifiedCount != 1 {
			log.Fatal("replace failed, expected 1 but got", result.MatchedCount)
		}

		if err = session.CommitTransaction(sc); err != nil {
			log.Fatal(err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	session.EndSession(ctx)

	var v bson.M
	if err = collection.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&v); err != nil {
		log.Fatal(err)
	}
	if v["year"] != int32(2000) {
	
		log.Fatal("expected 2000 but got", v["year"])
	}

	res, _ := collection.DeleteOne(ctx, bson.M{"_id": id})
	if res.DeletedCount != 1 {
		log.Fatal("delete failed, expected 1 but got", res.DeletedCount)
	}
}

func getMongoClient() (*mongo.Client, error) {
	uri := "mongodb://localhost:27017/argos?replicaSet=replset"
	if os.Getenv("DATABASE_URL") != "" {
		uri = os.Getenv("DATABASE_URL")
	}
	return getMongoClientByURI(uri)
}

func getMongoClientByURI(uri string) (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		return client, err
	}
	client.Ping(context.Background(), nil)
	return client, err
}

