package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/ratnakar5938/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://Ratnakar5938:<password>@cluster0.axkg3.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client option
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	checkErr(err)
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ready")
}

// MongoDB helpers - file

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	checkErr(err)
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkErr(err)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkErr(err)
	fmt.Println("Modified count:", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkErr(err)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	checkErr(err)
	fmt.Println("Movie got deleted with delete count:", deleteCount)
}

// delete all records from mongoDB
func deleteAllMovies() {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	checkErr(err)
	fmt.Println("DeleteCount:", deleteResult.DeletedCount)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
