package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	/** MONGO TESTING
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := "mongodb+srv://dbUser:dbUserPassword@cluster0-5nzly.mongodb.net/test?retryWrites=true&w=majority"

	client := connectToMongoDB(ctx, uri)
	feedbackCol := connectToCollection("classDB", "feedbackCollection", client)

	bsonData := bson.D{
		{"_id", 3},
		{"className", "sampleClass"},
		{"startTime", "12:00 PM"},
		{"professor", "sampleProfessor"},
		{"building", "Bascom Hall"},
		{"choice", "g"},
	}

	insertSingleResponse(ctx, feedbackCol, bsonData)
	fmt.Println(deleteSingleResponse(ctx, feedbackCol, 2))
	disconnectFromMongoCluster(ctx, client)
	err := client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("sucessful")
	}
	**/
}

/////////////////////////////////// MONGO METHODS ///////////////////////////////////

func connectToMongoDB(ctx context.Context, uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil
	}
	return client
}

func connectToCollection(databaseName string, collectionName string, client *mongo.Client) *mongo.Collection {
	classDB := client.Database(databaseName)
	return classDB.Collection(collectionName)
}

func disconnectFromMongoCluster(ctx context.Context, client *mongo.Client) bool {
	client.Disconnect(ctx)
	err := client.Ping(ctx, nil)
	if err == nil {
		return false
	}
	return true
}

func insertSingleResponse(ctx context.Context, collection *mongo.Collection, bsonData bson.D) bool {
	_, err := collection.InsertOne(ctx, bsonData)
	if err != nil {
		return false
	}
	return true
}

func deleteSingleResponse(ctx context.Context, collection *mongo.Collection, id int) bool {
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return false
	}
	return false
}
